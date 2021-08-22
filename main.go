package main

import (
	"github.com/brianvoe/gofakeit/v6"
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

	itemRelRepository := mysql.NewItemRepository(db)

	itemService := item.NewService(itemRelRepository)

	//result, err := itemService.GetByID(1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(result)
	//
	//result2, err := itemService.GetAll()
	//if err != nil {
	//	panic(err)
	//}

	//spew.Dump(result2)

	var newItem item.RelationalItem
	err = gofakeit.Struct(&newItem)
	if err != nil {
		panic(err)
	}

	id, err := itemService.Create(&newItem)
	if err != nil {
		panic(err)
	}

	spew.Dump(id)

}
