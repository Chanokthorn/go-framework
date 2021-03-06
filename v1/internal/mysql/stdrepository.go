package mysql

import (
	"fmt"
	"reflect"
	"reflect-test/v1/internal/mysql/model"
	"reflect-test/v1/internal/std"
	"reflect-test/v1/internal/std/mysql"
	"strconv"
	"strings"
)

type StdConfig struct {
	TableName         string
	IDField           string
	UUIDField         string
	ParentIDField     string
	RecursiveOnGetAll bool
}

func getConfig(t reflect.Type) (StdConfig, error) {
	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("std"); field != "" {
			config, err := parseConfig(field)
			if err != nil {
				return StdConfig{}, err
			}

			return config, nil
		}
	}

	return StdConfig{}, fmt.Errorf(`unable to find std config`)
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

	v.Elem().FieldByName("")
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

type DomainRepository struct {
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

	common := reflect.TypeOf(mysql.DBModelCommon{})

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

func newStandardRepository(obj interface{}, db *DB) (mysql.DomainRepository, error) {
	t := reflect.TypeOf(obj)

	config, err := getConfig(t)
	if err != nil {
		return nil, fmt.Errorf(`unable to get std config: %v`, err)
	}

	return &DomainRepository{t: t, config: config, db: db, fields: getFields(t)}, nil
}

func isAggregate(field interface{}) bool {
	inter := reflect.TypeOf((*std.DBAggregateModel)(nil)).Elem()

	if reflect.TypeOf(field).Implements(inter) {
		return true
	}

	return false
}

func implementsDBAggregateModel(t reflect.Type) bool {
	inter := reflect.TypeOf((*std.DBAggregateModel)(nil)).Elem()

	return reflect.PtrTo(t.Elem()).Implements(inter)
}

func implementsDBAggregateModelSlice(t reflect.Type) bool {
	if t.Kind() == reflect.Slice {
		return implementsDBAggregateModel(t)
	}

	return false
}

func (m *DomainRepository) GetAggregates(v reflect.Value, rootID int) error {
	t := v.Type().Elem()

	fields, config, err := getFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(fields, ", ") + ", CreatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + config.TableName)
	txtSQL.WriteString(" WHERE " + config.ParentIDField + " = ? AND IsDeleted = false")

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

func (m *DomainRepository) GetByID(id int) (std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " = ? AND IsDeleted = false")

	item := reflect.New(m.t).Interface()

	s := txtSQL.String()
	println(s)

	err := m.db.Get(item, txtSQL.String(), id)
	if err != nil {
		return nil, fmt.Errorf(`unable to get: %v`, err)
	}

	rootModel := item.(std.DBModel)

	v := reflect.ValueOf(rootModel)
	t := reflect.TypeOf(rootModel).Elem()

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.GetAggregates(v.Elem().Field(i), id)
			if err != nil {
				return nil, fmt.Errorf(`unable to get aggregates: %v`, err)
			}
		}
	}

	return v.Interface().(std.DBModel).ToDomain(), nil
}

func (m *DomainRepository) GetAll() ([]std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate, UpdatedBy, UpdatedDate")
	txtSQL.WriteString(" FROM " + m.config.TableName)

	items := reflect.New(reflect.SliceOf(m.t)).Interface()

	err := m.db.Select(items, txtSQL.String())
	if err != nil {
		return nil, fmt.Errorf(`unable to get all: %v`, err)
	}

	s := reflect.Indirect(reflect.ValueOf(items))

	result := []std.DomainModel{}

	for i := 0; i < s.Len(); i++ {
		result = append(result, s.Index(i).Interface().(model.DBItem).ToDomain())
	}

	return result, nil
}

func (m *DomainRepository) Search(domain std.DomainModel) ([]std.DomainModel, error) {
	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

	v := reflect.ValueOf(dbObject)
	t := reflect.TypeOf(dbObject).Elem()

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

	rows, err := m.db.NamedQuery(txtSQL.String(), dbObject)
	if err != nil {
		return nil, fmt.Errorf(`unable to get: %v`, err)
	}

	result := []std.DomainModel{}

	for rows.Next() {
		doc := reflect.New(m.t).Interface()

		err := rows.StructScan(doc)
		if err != nil {
			return nil, fmt.Errorf(`unable to scan struct: %v`, err)
		}

		result = append(result, doc.(std.DBModel).ToDomain())
	}

	return result, nil
}

// InsertAggregates does not support recursive insert in this implementation
func (m *DomainRepository) InsertAggregates(name string, v reflect.Value, rootID int) error {
	t := v.Type().Elem()

	fields, config, err := getFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	for i := 0; i < v.Len(); i++ {
		v.Index(i).FieldByName("ItemID").Set(reflect.ValueOf(&rootID))
	}

	err = setCreateFieldsSlice("system", v)
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

func (m *DomainRepository) Insert(domain std.DomainModel) (id int, err error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(m.config.TableName)
	txtSQL.WriteString("(" + strings.Join(m.fields, ", ") + ", CreatedBy, CreatedDate)")
	txtSQL.WriteString(" VALUES (:" + strings.Join(m.fields, ", :") + ", :CreatedBy, ADDDATE(NOW(), INTERVAL 7 HOUR))")

	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

	err = setCreateFields("system", reflect.ValueOf(dbObject))
	if err != nil {
		return 0, fmt.Errorf(`unable to set creaet fields: %v`, err)
	}

	res, err := m.db.NamedExec(txtSQL.String(), dbObject)
	if err != nil {
		return 0, fmt.Errorf(`unable to insert: %v`, err)
	}

	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf(`unable to get inserted id: %v`, err)
	}

	v := reflect.ValueOf(dbObject)
	t := reflect.TypeOf(dbObject).Elem()

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

// DeleteAggregates does not support recursive in this implementation
func (m *DomainRepository) DeleteAggregates(name string, t reflect.Type, rootID int) error {
	_, config, err := getFieldAndConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get config and field: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + config.TableName + " SET IsDeleted = true, UpdatedBy = ?, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")
	txtSQL.WriteString(" WHERE " + config.ParentIDField + " = " + strconv.Itoa(rootID))

	_, err = m.db.Exec(txtSQL.String(), name)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	return nil
}

func (m *DomainRepository) GetID(v reflect.Value) (int, error) {
	t := v.Type()

	config, err := getConfig(t)
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

func (m *DomainRepository) GetIDByUUID(uuid string) (int, error) {
	config, err := getConfig(m.t)
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

func (m *DomainRepository) Update(domain std.DomainModel) error {
	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

	err := setUpdateFields("system", reflect.ValueOf(dbObject))
	if err != nil {
		return fmt.Errorf(`unable to set create fields: %v`, err)
	}

	v := reflect.ValueOf(dbObject)
	t := reflect.TypeOf(dbObject).Elem()

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

	_, err = m.db.NamedExec(txtSQL.String(), dbObject)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	rootID, err := m.GetID(v.Elem())

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

func (m *DomainRepository) Delete(uuid string) error {
	rootID, err := m.GetIDByUUID(uuid)
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
