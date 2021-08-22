package mysql

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"reflect-test/mysql/model"
	"reflect-test/std"
	"strings"
)

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
	tableName string
	t         reflect.Type
	idField   string
	db        *DB
	fields    []string
}

func newStandardRepository(tableName string, obj interface{}, idField string, db *DB) std.Repository {
	t := reflect.TypeOf(obj)
	return &standardRepository{tableName: tableName, t: t, idField: idField, fields: getFields(t), db: db}
}

func (m *standardRepository) Get(id int) (std.DomainModel, error) {
	var txtSQL strings.Builder

	txtSQL.WriteString("SELECT ")
	txtSQL.WriteString(strings.Join(m.fields, ", "))
	txtSQL.WriteString(" FROM " + m.tableName)
	txtSQL.WriteString(" WHERE " + m.idField + " = ?")

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
	txtSQL.WriteString(" FROM " + m.tableName)

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
	txtSQL.WriteString(m.tableName)
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
