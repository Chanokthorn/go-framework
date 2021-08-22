package mysql

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"reflect-test/mysql/model"
	"reflect-test/std"
	"strings"
)

type StdConfig struct {
	TableName         string
	IDField           string
	UUIDField         string
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

func (m *standardRepository) GetByID(id int) (std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", "))
	txtSQL.WriteString(" FROM " + m.config.TableName)
	txtSQL.WriteString(" WHERE " + m.config.IDField + " = ?, AND IsDeleted = false")

	item := reflect.New(m.t).Interface()

	err := m.db.Get(item, txtSQL.String(), id)
	if err != nil {
		return nil, fmt.Errorf(`unable to get: %v`, err)
	}

	result := item.(std.DBModel)

	return result.ToModel(), nil
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
		result = append(result, s.Index(i).Interface().(model.DBItem).ToModel())
		john := s.Index(i).Interface().(model.DBItem)
		spew.Dump(john.ToModel())
	}

	return result, nil
}

func (m *standardRepository) Insert(domain std.DomainModel) (id int, err error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("INSERT INTO ")
	txtSQL.WriteString(m.config.TableName)
	txtSQL.WriteString("(" + strings.Join(m.fields, ", ") + ")")
	txtSQL.WriteString(" VALUES (:" + strings.Join(m.fields, ", :") + ")")

	name := txtSQL.String()
	println(name)

	dbObject := reflect.New(m.t).Interface().(std.DBModel)

	dbObject.Set(domain)

	res, err := m.db.NamedExec(txtSQL.String(), dbObject)
	if err != nil {
		return 0, fmt.Errorf(`unable to insert: %v`, err)
	}

	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf(`unable to get inserted id: %v`, err)
	}

	return int(id64), nil
}
func (m *standardRepository) Update(domain std.DomainModel) error {
	dbObject := reflect.New(m.t).Interface().(std.DBModel)

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

	str := txtSQL.String()
	println(str)

	res, err := m.db.NamedExec(txtSQL.String(), dbObject)
	if err != nil {
		return fmt.Errorf("unable to update: %v", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("unable to retrieve updated id: %v", err)
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
