package main

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"reflect-test/v2/internal/duck"
	"reflect-test/v2/internal/mysql"
	"reflect-test/v2/internal/std/mysql"
)

func main() {
	connect := "john:john@tcp(localhost:3307)/john?charset=utf8&allowOldPasswords=1&parseTime=true&loc=Asia%2FBangkok"

	db, err := std_mysql.NewMysqlDB(connect)
	if err != nil {
		panic(err)
	}

	duckRepository, err := mysql.NewDuckRepository(db)
	if err != nil {
		panic(err)
	}

	duckService := duck.NewService(duckRepository)

	ctx := context.TODO()
	spew.Dump()

	/// GET BY ID ///
	//d, err := duckService.GetByID(ctx, 29)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(d)

	/// GET BY UUID ///
	//d, err := duckService.GetByUUID(ctx, "23123462-f076-3017-89d4-635be9b90d6f")
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(d)

	/// GET ALL ///
	//ds, err := duckService.GetAll(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(ds)

	/// SEARCH ///
	//color := "white"
	//d := duck.Duck{Color: &color}
	//ds, err := duckService.Search(ctx, d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(ds)

	/// CREATE ///
	//var d duck.Duck
	//
	//err = gofakeit.Struct(&d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//id, err := duckService.Create(ctx, d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(id)

	/// UPDATE ///
	var d duck.Duck

	err = gofakeit.Struct(&d)
	if err != nil {
		panic(err)
	}

	uuid := "23123462-f076-3017-89d4-635be9b90d6f"
	d.DuckUUID = &uuid

	err = duckService.Update(ctx, d)
	if err != nil {
		panic(err)
	}

	/// DELETE ///
	//uuid := "96ae1799-8468-3a79-a332-666560aee516"
	//
	//err = duckService.Delete(ctx, uuid)
	//if err != nil {
	//	panic(err)
	//}

}
