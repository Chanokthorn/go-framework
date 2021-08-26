package main

import (
	"github.com/davecgh/go-spew/spew"
	"reflect-test/v1/internal/item"
	"reflect-test/v1/internal/mysql"
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

	////// SEARCH
	////uuid := "726d6cf7-4118-487d-9e2b-a3d7b0dc7e8d"
	//john := "john"
	//item := item.RelationalItem{
	//	//UUID: &uuid,
	//	Name: &john,
	//}
	//searchResult, err := itemService.Search(&item)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(searchResult)

	/////// DELETE ////////
	//err = itemService.Delete("17ab0e1b-87de-46ed-84ce-62453be249b5")
	//if err != nil {
	//	panic(err)
	//}

	//////// UPDATE //////
	//uuid := "17ab0e1b-87de-46ed-84ce-62453be249b5"
	//john := "updated name"
	//
	//var newItem item.RelationalItem
	//err = gofakeit.Struct(&newItem)
	//if err != nil {
	//	panic(err)
	//}
	//
	//newItem.UUID = &uuid
	//newItem.Name = &john
	//
	//err = itemService.Update(&newItem)
	//if err != nil {
	//	panic(err)
	//}

	///// GET BY ID ///////
	result, err := itemService.GetByID(150)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)

	//result2, err := itemService.GetAll()
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(result2)

	//////// CREATE /////
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
	//
	//spew.Dump(id)

}
