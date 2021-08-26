package std_mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
)

type Repository interface {
	GetByID(ctx context.Context, dest interface{}, id int) error
	GetAll(ctx context.Context, dest interface{}) error
	Search(ctx context.Context, dest interface{}, model DBModel) error
	Insert(ctx context.Context, model DBModel) (id int, err error)
	Update(ctx context.Context, model DBModel) error
	Delete(ctx context.Context, uuid string) error
}

func getConfig(t reflect.Type) (StdConfig, error) {
	return reflect.New(t).Interface().(DBModel).GetConfig(), nil
}

func getFields(t reflect.Type) []string {

	fields := []string{}

	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("db"); field != "" {
			fields = append(fields, field)
		}
	}

	return fields
}

func setCreateFields(name string, v reflect.Value) error {
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf(`value must be pointer`)
	}

	u, _ := uuid.FromString(uuid.NewV4().String())
	generatedUUID := uuid.NewV3(u, "www.qchang.com").String()

	v.Elem().FieldByName("UUID").Set(reflect.ValueOf(&generatedUUID))
	v.Elem().FieldByName("CreatedBy").Set(reflect.ValueOf(&name))

	return nil
}

func setCreateFieldsSlice(name string, v reflect.Value) error {
	if v.Kind() != reflect.Slice {
		return fmt.Errorf(`value must be slice`)
	}

	for i := 0; i < v.Len(); i++ {
		v.Index(i).FieldByName("CreatedBy").Set(reflect.ValueOf(&name))
	}

	return nil
}

func setUpdateFields(name string, v reflect.Value) error {
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf(`value must be pointer`)
	}

	v.Elem().FieldByName("UpdatedBy").Set(reflect.ValueOf(&name))

	return nil
}

func setUpdateFieldsSlice(name string, v reflect.Value) error {
	if v.Kind() != reflect.Slice {
		return fmt.Errorf(`value must be slice`)
	}

	for i := 0; i < v.Len(); i++ {
		v.Index(i).FieldByName("UpdatedBy").Set(reflect.ValueOf(&name))
	}

	return nil
}

type MysqlRepository struct {
	t      reflect.Type
	config StdConfig
	db     *DB
	fields []string
}

func validateDBModel(model interface{}) error {
	t := reflect.TypeOf(model).Elem()

	if _, err := getConfig(t); err != nil {
		return fmt.Errorf(`invalid model config: %v`, err)
	}

	common := reflect.TypeOf(DBModelCommon{})

	foundCommon := false
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == common {
			foundCommon = true
		}
	}

	if !foundCommon {
		return fmt.Errorf(`model must contain std.DBModelCommon`)
	}

	return nil
}

func getFieldAndConfig(t reflect.Type) ([]string, StdConfig, error) {
	config, err := getConfig(t)
	if err != nil {
		return nil, StdConfig{}, fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	return fields, config, nil
}

func NewRepository(obj interface{}, db *DB) (Repository, error) {
	t := reflect.TypeOf(obj)

	config, err := getConfig(t)
	if err != nil {
		return nil, fmt.Errorf(`unable to get std config: %v`, err)
	}

	return &MysqlRepository{t: t, config: config, db: db, fields: getFields(t)}, nil
}

func isAggregate(field interface{}) bool {
	inter := reflect.TypeOf((*DBAggregateModel)(nil)).Elem()

	if reflect.TypeOf(field).Implements(inter) {
		return true
	}

	return false
}

func implementsDBAggregateModel(t reflect.Type) bool {
	inter := reflect.TypeOf((*DBAggregateModel)(nil)).Elem()

	return reflect.PtrTo(t.Elem()).Implements(inter)
}

func implementsDBAggregateModelSlice(t reflect.Type) bool {
	if t.Kind() == reflect.Slice {
		return implementsDBAggregateModel(t)
	}

	return false
}

func (m *MysqlRepository) GetByID(ctx context.Context, dest interface{}, id int) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " = ? AND IsDeleted = false")

	s := txtSQL.String()

	println(s)

	err := m.db.Get(dest, txtSQL.String(), id)
	if err != nil {
		return fmt.Errorf(`unable to get: %v`, err)
	}

	return nil
}

func (m *MysqlRepository) GetAll(ctx context.Context, dest interface{}) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	err := m.db.Select(items, txtSQL.String())
	if err != nil {
		return fmt.Errorf(`unable to get all: %v`, err)
	}

	s := reflect.Indirect(reflect.ValueOf(items))

	v := reflect.ValueOf(dest).Elem()

	for i := 0; i < s.Len(); i++ {
		v.Set(reflect.Append(v, s.Index(i)))
	}

	return nil
}

func (m *MysqlRepository) Search(ctx context.Context, dest interface{}, model DBModel) error {
	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model).Elem()

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE IsDeleted = false")

	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("db"); field != "" {
			if !v.Elem().Field(i).IsNil() {
				txtSQL.WriteString(" AND " + field + " = :" + field)
			}
		}
	}

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	query, args, err := sqlx.Named(txtSQL.String(), model)
	query, args, err = sqlx.In(query, args...)
	query = m.db.Rebind(query)

	err = m.db.Select(items, query, args...)
	if err != nil {
		return fmt.Errorf(`unable to get all: %v`, err)
	}

	s := reflect.Indirect(reflect.ValueOf(items))

	destV := reflect.ValueOf(dest).Elem()

	for i := 0; i < s.Len(); i++ {
		destV.Set(reflect.Append(destV, s.Index(i)))
	}

	return nil
}

func (m *MysqlRepository) Insert(ctx context.Context, model DBModel) (id int, err error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(m.config.TableName)
	txtSQL.WriteString("(" + strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate)")
	txtSQL.WriteString(" VALUES (:" + strings.Join(m.fields, ", :") + ", :CreatedBy, ADDDATE(NOW(), INTERVAL 7 HOUR))")

	err = setCreateFields("system", reflect.ValueOf(model))
	if err != nil {
		return 0, fmt.Errorf(`unable to set creaet fields: %v`, err)
	}

	res, err := m.db.NamedExec(txtSQL.String(), model)
	if err != nil {
		return 0, fmt.Errorf(`unable to insert: %v`, err)
	}

	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf(`unable to get inserted id: %v`, err)
	}

	return int(id64), nil
}

func (m *MysqlRepository) Update(ctx context.Context, model DBModel) error {
	panic("implement me")
}

func (m *MysqlRepository) Delete(ctx context.Context, uuid string) error {
	panic("implement me")
}
