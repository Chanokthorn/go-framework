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
	GetByIDs(ctx context.Context, dest interface{}, ids []int) error
	GetByUUIDs(ctx context.Context, dest interface{}, uuids []string) error
	GetByID(ctx context.Context, dest interface{}, id int) error
	GetByUUID(ctx context.Context, dest interface{}, uuid string) error
	FillStructsByID(ctx context.Context, src interface{}) error
	FillStructsByUUID(ctx context.Context, src interface{}) error
	FillStructByID(ctx context.Context, src interface{}) error
	FillStructByUUID(ctx context.Context, src interface{}) error
	GetByRootID(ctx context.Context, dest interface{}, rootID int) error
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

func getFields(t reflect.Type) []string {

	fields := []string{}

	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("db"); field != "" {
			fields = append(fields, field)
		}
	}

	return fields
}

func setCreateFields(name string, config RootModelConfig, v reflect.Value) error {
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf(`value must be pointer`)
	}

	u, _ := uuid.FromString(uuid.NewV4().String())
	generatedUUID := uuid.NewV3(u, "www.qchang.com").String()

	v.Elem().FieldByName(config.UUIDField).Set(reflect.ValueOf(&generatedUUID))
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

func (m *MysqlRepository) getAggregates(v reflect.Value, rootID int) error {
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

func generateIntSliceString(intSlice []int, prefix string, delim string) string {
	var txt strings.Builder

	for i, int := range intSlice {
		if i > 0 {
			txt.WriteString(delim)
		}

		txt.WriteString(prefix + strconv.Itoa(int))
	}

	return txt.String()
}

func (m *MysqlRepository) GetByIDs(ctx context.Context, dest interface{}, ids []int) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " IN (" + generateIntSliceString(ids, "", ", ") + ") AND IsDeleted = false")
	txtSQL.WriteString(" ORDER BY FIELD(" + m.config.IDField + ", " + generateIntSliceString(ids, "", ", ") + ")")

	ss := txtSQL.String()

	println(ss)

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

func (m *MysqlRepository) GetByUUIDs(ctx context.Context, dest interface{}, uuids []string) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " IN ('" + strings.Join(uuids, "', '") + "') AND IsDeleted = false")
	txtSQL.WriteString(" ORDER BY FIELD(" + m.config.UUIDField + ", '" + strings.Join(uuids, "', '") + "')")

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

func (m *MysqlRepository) GetByID(ctx context.Context, dest interface{}, id int) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " = ? AND IsDeleted = false")

	err := m.db.Get(dest, txtSQL.String(), id)
	if err != nil {
		return fmt.Errorf(`unable to get: %v`, err)
	}

	v := reflect.ValueOf(dest)
	t := reflect.TypeOf(dest).Elem()

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.getAggregates(v.Elem().Field(i), id)
			if err != nil {
				return fmt.Errorf(`unable to get aggregates: %v`, err)
			}
		}
	}

	return nil
}

func (m *MysqlRepository) GetByUUID(ctx context.Context, dest interface{}, uuid string) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " = ? AND IsDeleted = false")

	err := m.db.Get(dest, txtSQL.String(), uuid)
	if err != nil {
		return fmt.Errorf(`unable to get: %v`, err)
	}

	v := reflect.ValueOf(dest)
	t := reflect.TypeOf(dest).Elem()

	id := int(v.Elem().FieldByName(m.config.IDField).Elem().Int())
	if id == 0 {
		return fmt.Errorf(`unable to get id`)
	}

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.getAggregates(v.Elem().Field(i), id)
			if err != nil {
				return fmt.Errorf(`unable to get aggregates: %v`, err)
			}
		}
	}

	return nil
}

func (m *MysqlRepository) FillStructsByID(ctx context.Context, src interface{}) error {
	v := reflect.ValueOf(src)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf(`src must be pointer to slice`)
	}

	if reflect.TypeOf(src).Elem().Elem() != m.t {
		fmt.Errorf(`invalid input type, must be %s`, m.t.String())
	}

	ids := []int{}
	for i := 0; i < v.Elem().Len(); i++ {
		ids = append(ids, int(v.Elem().Index(i).FieldByName(m.config.IDField).Elem().Int()))
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " IN (" + generateIntSliceString(ids, "", ", ") + ") AND IsDeleted = false")
	txtSQL.WriteString(" ORDER BY FIELD(" + m.config.IDField + ", " + generateIntSliceString(ids, "", ", ") + ")")

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	err := m.db.Select(items, txtSQL.String())
	if err != nil {
		return fmt.Errorf(`unable to get all: %v`, err)
	}

	v.Elem().Set(reflect.ValueOf(items).Elem())

	return nil
}

func (m *MysqlRepository) FillStructsByUUID(ctx context.Context, src interface{}) error {
	v := reflect.ValueOf(src)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf(`src must be pointer to slice`)
	}

	if reflect.TypeOf(src).Elem().Elem() != m.t {
		fmt.Errorf(`invalid input type, must be %s`, m.t.String())
	}

	uuids := []string{}
	for i := 0; i < v.Elem().Len(); i++ {
		uuids = append(uuids, v.Elem().Index(i).FieldByName(m.config.UUIDField).Elem().String())
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " IN ('" + strings.Join(uuids, "', '") + "') AND IsDeleted = false")
	txtSQL.WriteString(" ORDER BY FIELD(" + m.config.UUIDField + ", '" + strings.Join(uuids, "', '") + "')")

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	err := m.db.Select(items, txtSQL.String())
	if err != nil {
		return fmt.Errorf(`unable to get all: %v`, err)
	}

	v.Elem().Set(reflect.ValueOf(items).Elem())

	return nil
}

func (m *MysqlRepository) FillStructByID(ctx context.Context, src interface{}) error {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf(`src must be pointer to struct`)
	}

	if v.Elem().Type() != m.t {
		return fmt.Errorf(`invalid input type, must be %s`, m.t.String())
	}

	id64 := v.Elem().FieldByName(m.config.IDField).Elem().Int()

	return m.GetByID(ctx, src, int(id64))
}

func (m *MysqlRepository) FillStructByUUID(ctx context.Context, src interface{}) error {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf(`src must be pointer to struct`)
	}

	if v.Elem().Type() != m.t {
		return fmt.Errorf(`invalid input type, must be %s`, m.t.String())
	}

	uuid := v.Elem().FieldByName(m.config.UUIDField).Elem().String()

	return m.GetByUUID(ctx, src, uuid)
}

func (m *MysqlRepository) GetByRootID(ctx context.Context, dest interface{}, rootID int) error {
	if m.config.RootIDField == "" {
		return fmt.Errorf(`the type does not have root id field`)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.RootIDField + " = ? AND IsDeleted = false")

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	err := m.db.Select(items, txtSQL.String(), rootID)
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
		v.Index(i).FieldByName(config.RootIDField).Set(reflect.ValueOf(&rootID))
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

	err = setCreateFields("system", m.config, reflect.ValueOf(model))
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

	uuid := v.FieldByName(m.config.UUIDField).Elem().Interface().(string)
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
