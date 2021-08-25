package mysql

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"reflect-test/mysql/model"
	"reflect-test/std"
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

type standardRepository struct {
	t      reflect.Type
	config StdConfig
	db     *DB
	fields []string
}

func newStandardRepository(obj interface{}, db *DB) (std.Repository, error) {
	t := reflect.TypeOf(obj)

	config, err := getConfig(t)
	if err != nil {
		return nil, fmt.Errorf(`unable to get std config: %v`, err)
	}

	return &standardRepository{t: t, config: config, db: db, fields: getFields(t)}, nil
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

func (m *standardRepository) GetAggregates(v reflect.Value, rootID int) error {
	t := v.Type().Elem()
	spew.Dump(t)

	config, err := getConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(fields, ", "))
	txtSQL.WriteString(" FROM " + config.TableName)
	txtSQL.WriteString(" WHERE " + config.ParentIDField + " = ? AND IsDeleted = false")

	s2 := txtSQL.String()

	println(s2)

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

func (m *standardRepository) GetByID(id int) (std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", "))
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

func (m *standardRepository) GetAll() ([]std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", "))
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
		john := s.Index(i).Interface().(model.DBItem)
		spew.Dump(john.ToDomain())
	}

	return result, nil
}

func (m *standardRepository) Search(domain std.DomainModel) ([]std.DomainModel, error) {
	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

	v := reflect.ValueOf(dbObject)
	t := reflect.TypeOf(dbObject).Elem()

	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", "))
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
func (m *standardRepository) InsertAggregates(v reflect.Value, rootID int) error {
	t := v.Type().Elem()
	spew.Dump(t)

	config, err := getConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get type config: %v`, err)
	}

	fields := getFields(t)

	for i := 0; i < v.Len(); i++ {
		spew.Dump(v.Index(i).Interface().(model.DBLocation))

		v.Index(i).FieldByName("ItemID").Set(reflect.ValueOf(&rootID))
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(config.TableName)
	txtSQL.WriteString("(" + strings.Join(fields, ", ") + ")")
	txtSQL.WriteString(" VALUES (:" + strings.Join(fields, ", :") + ")")

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

func (m *standardRepository) Insert(domain std.DomainModel) (id int, err error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(m.config.TableName)
	txtSQL.WriteString("(" + strings.Join(m.fields, ", ") + ")")
	txtSQL.WriteString(" VALUES (:" + strings.Join(m.fields, ", :") + ")")

	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

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
			err = m.InsertAggregates(v.Elem().Field(i), int(id64))
			if err != nil {
				return 0, fmt.Errorf(`unable to insert aggregates: %v`, err)
			}
		}
	}

	return int(id64), nil
}

// DeleteAggregates does not support recursive in this implementation
func (m *standardRepository) DeleteAggregates(t reflect.Type, rootID int) error {
	config, err := getConfig(t)
	if err != nil {
		return fmt.Errorf(`unable to get type config: %v`, err)
	}

	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + config.TableName + " SET IsDeleted = true, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")
	txtSQL.WriteString(" WHERE " + config.ParentIDField + " = " + strconv.Itoa(rootID))

	_, err = m.db.Exec(txtSQL.String())
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	return nil
}

func (m *standardRepository) GetID(v reflect.Value) (int, error) {
	t := v.Type()

	spew.Dump(t)

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

func (m *standardRepository) Update(domain std.DomainModel) error {
	dbObject := reflect.New(m.t).Interface().(std.DBRootModel)

	dbObject.Set(domain)

	v := reflect.ValueOf(dbObject)
	t := reflect.TypeOf(dbObject).Elem()

	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + m.config.TableName + " SET UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")

	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i).Tag.Get("db"); field != "" {
			if !v.Elem().Field(i).IsNil() {
				txtSQL.WriteString(", " + field + " = :" + field)
			}
		}
	}

	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " = :" + m.config.UUIDField)

	_, err := m.db.NamedExec(txtSQL.String(), dbObject)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	rootID, err := m.GetID(v.Elem())

	for i := 0; i < t.NumField(); i++ {
		if implementsDBAggregateModelSlice(t.Field(i).Type) {
			err = m.DeleteAggregates(t.Field(i).Type.Elem(), rootID)
			if err != nil {
				return fmt.Errorf(`unable to delete aggregates: %v`, err)
			}

			err = m.InsertAggregates(v.Elem().Field(i), rootID)
			if err != nil {
				return fmt.Errorf(`unable to insert aggregates: %v`, err)
			}
		}
	}

	return nil

}

func (m *standardRepository) Delete(uuid string) error {
	var txtSQL strings.Builder

	txtSQL.WriteString("UPDATE " + m.config.TableName + " SET IsDeleted = true, UpdatedDate = ADDDATE(NOW(), INTERVAL 7 HOUR)")
	txtSQL.WriteString(" WHERE " + m.config.UUIDField + " = " + uuid)

	res, err := m.db.Exec(txtSQL.String())
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

	return nil
}
