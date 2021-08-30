package std_mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strconv"
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

func getRootConfig(t reflect.Type) (RootModelConfig, error) {
	return reflect.New(t).Interface().(DBRootModel).GetConfig(), nil
}

func getAggregateConfig(t reflect.Type) (AggregateModelConfig, error) {
	return reflect.New(t).Interface().(DBAggregateModel).GetConfig(), nil
}

func getUpdAggregateConfig(t reflect.Type) (UpdatableAggregateModelConfig, error) {
	return reflect.New(t).Interface().(DBUpdatableAggregateModel).GetConfig(), nil
}

func getRootFieldAndConfig(t reflect.Type) ([]string, RootModelConfig, error) {
	config, err := getRootConfig(t)
	if err != nil {
		return nil, RootModelConfig{}, fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	return fields, config, nil
}

func getAggregateFieldAndConfig(t reflect.Type) ([]string, AggregateModelConfig, error) {
	config, err := getAggregateConfig(t)
	if err != nil {
		return nil, AggregateModelConfig{}, fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	return fields, config, nil
}

func getUpdAggregateFieldAndConfig(t reflect.Type) ([]string, UpdatableAggregateModelConfig, error) {
	config, err := getUpdAggregateConfig(t)
	if err != nil {
		return nil, UpdatableAggregateModelConfig{}, fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	return fields, config, nil
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
	config RootModelConfig
	db     *DB
	fields []string
}

func validateDBModel(model interface{}) error {
	t := reflect.TypeOf(model).Elem()

	if _, err := getRootConfig(t); err != nil {
		return fmt.Errorf(`invalid model config: %v`, err)
	}

	common := reflect.TypeOf(DBRootCommon{})

	foundCommon := false
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == common {
			foundCommon = true
		}
	}

	if !foundCommon {
		return fmt.Errorf(`model must contain std.DBRootCommon`)
	}

	return nil
}

func NewRepository(obj interface{}, db *DB) (Repository, error) {
	t := reflect.TypeOf(obj)

	config, err := getRootConfig(t)
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

func (m *MysqlRepository) GetAggregates(v reflect.Value, rootID int) error {
	t := v.Type().Elem()

	fields, config, err := getAggregateFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(fields, ", ") + ", CreatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + config.TableName)
	txtSQL.WriteString(" WHERE " + config.RootIDField + " = ? AND IsDeleted = false")

	items := reflect.New(reflect.SliceOf(t)).Interface()

	err = m.db.Select(items, txtSQL.String(), rootID)
	if err != nil {
		return fmt.Errorf(`unable to get all: %v`, err)
	}

	s := reflect.Indirect(reflect.ValueOf(items))

	for i := 0; i < s.Len(); i++ {
		v.Set(reflect.Append(v, s.Index(i)))
	}

	return nil
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

	v := reflect.ValueOf(dest)
	t := reflect.TypeOf(dest).Elem()

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.GetAggregates(v.Elem().Field(i), id)
			if err != nil {
				return fmt.Errorf(`unable to get aggregates: %v`, err)
			}
		}
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

// InsertAggregates does not support recursive insert in this implementation
func (m *MysqlRepository) InsertAggregates(name string, v reflect.Value, rootID int) error {
	t := v.Type().Elem()

	fields, config, err := getAggregateFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	for i := 0; i < v.Len(); i++ {
		v.Index(i).FieldByName("RootID").Set(reflect.ValueOf(&rootID))
	}

	err = setCreateFieldsSlice(name, v)
	if err != nil {
		return fmt.Errorf(`unable to set create fields: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(config.TableName)
	txtSQL.WriteString("(" + strings.Join(fields, ", ") + ", CreatedBy, CreatedDate)")
	txtSQL.WriteString(" VALUES (:" + strings.Join(fields, ", :") + ", :CreatedBy, ADDDATE(NOW(), INTERVAL 7 HOUR))")

	res, err := m.db.NamedExec(txtSQL.String(), v.Interface())
	if err != nil {
		return fmt.Errorf(`unable to insert: %v`, err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf(`unable to inserted id: %v`, err)
	}

	// TODO: recursive implementation

	return nil
}

// InsertAggregates does not support recursive insert in this implementation
func (m *MysqlRepository) InsertUpdAggregates(name string, v reflect.Value, rootID int) error {
	t := v.Type().Elem()

	fields, config, err := getUpdAggregateFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	for i := 0; i < v.Len(); i++ {
		v.Index(i).FieldByName("RootID").Set(reflect.ValueOf(&rootID))
	}

	err = setCreateFieldsSlice(name, v)
	if err != nil {
		return fmt.Errorf(`unable to set create fields: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(config.TableName)
	txtSQL.WriteString("(" + strings.Join(fields, ", ") + ", CreatedBy, CreatedDate)")
	txtSQL.WriteString(" VALUES (:" + strings.Join(fields, ", :") + ", :CreatedBy, ADDDATE(NOW(), INTERVAL 7 HOUR))")

	res, err := m.db.NamedExec(txtSQL.String(), v.Interface())
	if err != nil {
		return fmt.Errorf(`unable to insert: %v`, err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf(`unable to inserted id: %v`, err)
	}

	// TODO: recursive implementation

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

	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model).Elem()

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.InsertAggregates("system", v.Elem().Field(i), int(id64))
			if err != nil {
				return 0, fmt.Errorf(`unable to insert aggregates: %v`, err)
			}
		}
	}

	return int(id64), nil
}

func (m *MysqlRepository) GetID(v reflect.Value) (int, error) {
	t := v.Type()

	config, err := getRootConfig(t)
	if err != nil {
		return 0, fmt.Errorf(`unable to get type config: %v`, err)
	}

	uuid := v.FieldByName("UUID").Elem().Interface().(string)
	if uuid == "" {
		return 0, fmt.Errorf(`uuid is emtpy`)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT " + config.IDField + " AS id")
	txtSQL.WriteString(" FROM " + config.TableName)
	txtSQL.WriteString(" WHERE " + config.UUIDField + " = ? AND IsDeleted = false")

	var id int

	err = m.db.Get(&id, txtSQL.String(), uuid)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteAggregates does not support recursive in this implementation
func (m *MysqlRepository) DeleteAggregates(name string, t reflect.Type, rootID int) error {
	_, config, err := getAggregateFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + config.TableName + " SET IsDeleted = true, UpdatedBy = ?, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")
	txtSQL.WriteString(" WHERE " + config.RootIDField + " = " + strconv.Itoa(rootID))

	_, err = m.db.Exec(txtSQL.String(), name)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	return nil
}

func (m *MysqlRepository) Update(ctx context.Context, model DBModel) error {
	err := setUpdateFields("system", reflect.ValueOf(model))
	if err != nil {
		return fmt.Errorf(`unable to set create fields: %v`, err)
	}

	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model).Elem()

	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + m.config.TableName + " SET UpdatedBy = :UpdatedBy, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")

	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("db"); field != "" {
			if !v.Elem().Field(i).IsNil() {
				txtSQL.WriteString(", " + field + " = :" + field)
			}
		}
	}

	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " = :" + m.config.UUIDField)

	_, err = m.db.NamedExec(txtSQL.String(), model)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	rootID, err := m.GetID(v.Elem())
	if err != nil {
		return fmt.Errorf("unable to get root id: %v", err)
	}

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.DeleteAggregates("system", t.Field(i).Type.Elem(), rootID)
			if err != nil {
				return fmt.Errorf(`unable to delete aggregates: %v`, err)
			}

			err = m.InsertAggregates("system", v.Elem().Field(i), rootID)
			if err != nil {
				return fmt.Errorf(`unable to insert aggregates: %v`, err)
			}
		}
	}

	return nil
}

func (m *MysqlRepository) GetRootIDByUUID(uuid string) (int, error) {
	config, err := getRootConfig(m.t)
	if err != nil {
		return 0, fmt.Errorf(`unable to get type config: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT " + config.IDField + " AS id")
	txtSQL.WriteString(" FROM " + config.TableName)
	txtSQL.WriteString(" WHERE " + config.UUIDField + " = ? AND IsDeleted = false")

	var id int

	err = m.db.Get(&id, txtSQL.String(), uuid)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *MysqlRepository) Delete(ctx context.Context, uuid string) error {
	rootID, err := m.GetRootIDByUUID(uuid)
	if err != nil {
		return fmt.Errorf(`unable to get id: %v`, err)
	}

	var txtSQL strings.Builder

	name := "system"

	txtSQL.WriteString("UPDATE " + m.config.TableName + " SET IsDeleted = true, UpdatedBy = ?, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")
	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " = '" + uuid + "'")

	res, err := m.db.Exec(txtSQL.String(), name)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to retrieve affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affeccted")
	}

	t := m.t

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.DeleteAggregates("system", t.Field(i).Type.Elem(), rootID)
			if err != nil {
				return fmt.Errorf(`unable to delete aggregates: %v`, err)
			}
		}
	}

	return nil
}
