package main

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"reflect-test/v3/internal/duck"
	std_mysql "reflect-test/v3/internal/lib/mysql"
	"reflect-test/v3/internal/lib/std"
	"reflect-test/v3/internal/lib/user"
	"reflect-test/v3/internal/mysql"
	"reflect-test/v3/internal/mysql/model"
	"reflect-test/v3/internal/pond"
)

func main() {
	connect := "john:john@tcp(localhost:3307)/john?charset=utf8&allowOldPasswords=1&parseTime=true&loc=Asia%2FBangkok"

	db, err := std_mysql.NewMysqlDB(connect)
	if err != nil {
		panic(err)
	}

	std.SetTxProvider(db, nil)

	pondRepository, err := std_mysql.NewRepository(pond.Pond{}, db)
	if err != nil {
		panic(err)
	}

	duckDBRepository, err := std_mysql.NewRepository(model.Duck{}, db)
	if err != nil {
		panic(err)
	}

	duckRepository, err := mysql.NewDuckRepository(db)
	if err != nil {
		panic(err)
	}

	duckService := duck.NewService(duckRepository)

	ctx := context.TODO()
	println(duckDBRepository, duckRepository, duckService)
	spew.Dump()
	gofakeit.New(1)
	profile := std.Profile{user.UserPermission{UserUUID: "1212312121"}}

	/// GET BY ID SLICE ///
	//err = std.WithProfile(ctx, profile, func(ctx context.Context) error {
	//	ds, err := duckRepository.GetByIDs(ctx, []int{28, 27, 29})
	//	if err != nil {
	//		return err
	//	}
	//
	//	spew.Dump(ds)
	//	return nil
	//})
	//
	//if err != nil {
	//	panic(err)
	//}

	/// GET BY UUID SLICE ///
	//ds, err := duckRepository.GetByUUIDs(ctx, []string{"dbb90ee9-8d45-340c-8f28-9496a7f3aefe", "35698f21-32dd-37a6-8828-a483dec40c13", "23123462-f076-3017-89d4-635be9b90d6f"})
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(ds)

	/// FILL STRUCT BY ID ///
	//var d model.Duck
	//id := 27
	//d.DuckID = &id
	//err = duckDBRepository.FillStructByID(ctx, &d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(d)

	/// FILL STRUCT BY UUID ///
	//var d model.Duck
	//uuid := "35698f21-32dd-37a6-8828-a483dec40c13"
	//d.DuckUUID = &uuid
	//err = duckDBRepository.FillStructByUUID(ctx, &d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(d)

	/// FILL STRUCTS BY ID ///
	//id1 := 28
	//id2 := 27
	//id3 := 29
	//ds := []model.Duck{{DuckID: &id1}, {DuckID: &id2}, {DuckID: &id3}}
	//err = duckDBRepository.FillStructsByID(ctx, &ds)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(ds)

	/// FILL STRUCTS BY ID ///
	//uuid1 := "dbb90ee9-8d45-340c-8f28-9496a7f3aefe"
	//uuid2 := "35698f21-32dd-37a6-8828-a483dec40c13"
	//uuid3 := "23123462-f076-3017-89d4-635be9b90d6f"
	//ds := []model.Duck{{DuckUUID: &uuid1}, {DuckUUID: &uuid2}, {DuckUUID: &uuid3}}
	//err = duckDBRepository.FillStructsByUUID(ctx, &ds)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(ds)

	/// GET BY ID ///
	//d, err := duckService.GetByID(ctx, 29)
	//if err != nil {
	//	panic(err)
	//}
	//
	//spew.Dump(d)

	/// GET BY UUID ///
	//d, err := duckService.GetByUUID(ctx, "1d51839e-b417-3680-9897-a7a3f0e008f9")
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
	//std.WithProfile(ctx, profile, func(ctx context.Context) error {
	//	err = std.WithRelTxContext(ctx, func(ctx context.Context) error {
	//		id, err := duckService.Create(ctx, d)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		spew.Dump(id)
	//
	//		return nil
	//	})
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return nil
	//})

	/// UPDATE ///
	//var d duck.Duck
	//
	//err = gofakeit.Struct(&d)
	//if err != nil {
	//	panic(err)
	//}
	//
	//uuid := "33c37a0c-2c7f-3862-89bf-63bacd2f499a"
	//d.DuckUUID = &uuid
	//
	//std.WithRelTxContext(ctx, func(ctx context.Context) error {
	//	err = duckService.Update(ctx, d)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return nil
	//})

	/// DELETE ///
	//uuid := "33c37a0c-2c7f-3862-89bf-63bacd2f499a"
	//
	//std.WithRelTxContext(ctx, func(ctx context.Context) error {
	//	err = duckService.Delete(ctx, uuid)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return nil
	//})

	///////////// POND //////////////
	var p pond.Pond

	err = gofakeit.Struct(&p)
	if err != nil {
		panic(err)
	}

	std.WithProfile(ctx, profile, func(ctx context.Context) error {
		err = std.WithRelTxContext(ctx, func(ctx context.Context) error {
			id, err := pondRepository.Insert(ctx, &p)
			if err != nil {
				panic(err)
			}

			spew.Dump(id)

			return nil
		})

		if err != nil {
			panic(err)
		}

		return nil
	})

}
