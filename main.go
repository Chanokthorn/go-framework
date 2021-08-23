package main

import (
	"github.com/davecgh/go-spew/spew"
	"reflect-test/item"
	"reflect-test/mysql"
)

func main() {
	connect := "john:john@tcp(localhost:3307)/john?charset=utf8&allowOldPasswords=1&parseTime=true&loc=Asia%2FBangkok"

	db, err := mysql.NewDB(connect)
	if err != nil {
		panic(err)
	}

	itemRelRepository, err := mysql.NewItemRepository(db)
	if err != nil {
		panic(err)
	}

	itemService := item.NewService(itemRelRepository)

	err = itemService.Delete("585647342")
	if err != nil {
		panic(err)
	}

	//uuid := "726d6cf7-4118-487d-9e2b-a3d7b0dc7e8d"
	//john := "john"
	//newItem := item.Item{
	//	UUID: &uuid,
	//	Name: &john,
	//}
	//
	//err = itemService.Update(&newItem)
	//if err != nil {
	//	panic(err)
	//}

	//result, err := itemService.GetByID(1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(result)
	//
	result2, err := itemService.GetAll()
	if err != nil {
		panic(err)
	}

	spew.Dump(result2)

	//var newItem item.RelationalItem
	//err = gofakeit.Struct(&newItem)
	//if err != nil {
	//	panic(err)
	//}
	//
	//id, err := itemService.Create(&newItem)
	//if err != nil {
	//	panic(err)
	//}

	//spew.Dump(id)

}
