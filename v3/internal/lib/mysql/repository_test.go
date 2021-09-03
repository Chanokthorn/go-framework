// docker-compose must be run before starting this test
package std_mysql

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"reflect-test/v3/internal/lib/std"
	"reflect-test/v3/internal/lib/user"
	"testing"
)

var repository Repository

const (
	connectionString = "test:test@tcp(localhost:3308)/test?charset=utf8&allowOldPasswords=1&parseTime=true&loc=Asia%2FBangkok&multiStatements=true"
	migrate          = `
	create table if not exists duck
	(
		DuckID      int auto_increment primary key,
		DuckUUID    text             not null,
		Name        text             null,
		Color       text             null,
		IsActive    bit default b'1' null,
		CreatedBy   text             null,
		CreatedDate datetime         null,
		UpdatedBy   text             null,
		UpdatedDate datetime         null,
		IsDeleted   bit default b'0' null
	);

	create table if not exists egg
	(
		EggID       int auto_increment primary key,
		DuckID      int              not null,
		Name        text             null,
		Age         int              null,
		IsActive    bit default b'1' null,
		CreatedBy   text             null,
		CreatedDate datetime         null,
		UpdatedBy   text             null,
		UpdatedDate datetime         null,
		IsDeleted   bit default b'0' null
	);
	`
	teardown = `
		drop table if exists duck;
		drop table if exists egg;
	`
	populate = `
	INSERT INTO test.duck (DuckID, DuckUUID, Name, Color, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, '0020c585-dfd2-3696-8245-6a211cb694a8', 'Conner Altenwerth', 'Teal', true, '123123123', '2021-09-02 15:14:31', null, null, false);
	INSERT INTO test.duck (DuckID, DuckUUID, Name, Color, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, '6fac5a9b-f3e4-35db-8236-7ead28ee7906', 'Al Will', 'OldLace', true, '123123123', '2021-09-02 15:14:49', null, null, false);
	INSERT INTO test.duck (DuckID, DuckUUID, Name, Color, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, '1e1f0ffd-96ed-3c07-b73d-5d172ca2380b', 'Olga Daugherty', 'SkyBlue', true, '123123123', '2021-09-02 15:14:51', null, null, false);
	INSERT INTO test.duck (DuckID, DuckUUID, Name, Color, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, 'e5a4c708-81b6-3a98-9fcb-5ff72c5394ef', 'Clovis Spinka', 'LightCyan', true, '123123123', '2021-09-02 15:14:53', null, null, false);

	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, 1, 'Lorine Nikolaus', 29, true, '123123123', '2021-09-02 15:14:31', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, 1, 'Garrison White', 2, true, '123123123', '2021-09-02 15:14:31', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, 1, 'Reta Skiles', 35, true, '123123123', '2021-09-02 15:14:31', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, 2, 'Adelle Connelly', 2, true, '123123123', '2021-09-02 15:14:49', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (5, 2, 'Dashawn Metz', 33, true, '123123123', '2021-09-02 15:14:49', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (6, 2, 'Keon Howe', 36, true, '123123123', '2021-09-02 15:14:49', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (7, 3, 'Maurine Yundt', 12, true, '123123123', '2021-09-02 15:14:51', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (8, 3, 'Karley Kozey', 7, true, '123123123', '2021-09-02 15:14:51', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (9, 3, 'Jermey Miller', 16, true, '123123123', '2021-09-02 15:14:51', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (10, 4, 'Buddy Cummerata', 37, true, '123123123', '2021-09-02 15:14:53', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (11, 4, 'Dustin Klocko', 20, true, '123123123', '2021-09-02 15:14:53', null, null, false);
	INSERT INTO test.egg (EggID, DuckID, Name, Age, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (12, 4, 'Alisha Hintz', 43, true, '123123123', '2021-09-02 15:14:53', null, null, false);
	`
)

type Duck struct {
	RootCommon
	DuckID   *int    `db:"DuckID" fake:"skip"`
	DuckUUID *string `db:"DuckUUID" fake:"skip"`
	Name     *string `db:"Name" fake:"{name}"`
	Color    *string `db:"Color" fake:"{color}"`
	Eggs     []Egg   `fakesize:"3"`
}

func (d *Duck) GetConfig() RootConfig {
	return RootConfig{
		TableName: "duck",
		IDField:   "DuckID",
		UUIDField: "DuckUUID",
	}
}

type Egg struct {
	AggregateCommon
	EggID  *int    `db:"EggID" fake:"skip"`
	DuckID *int    `db:"DuckID" fake:"skip"`
	Name   *string `db:"Name" fake:"{name}"`
	Age    *int    `db:"Age" fake:"{number:1,50}"`
}

func (e *Egg) GetConfig() AggregateConfig {
	return AggregateConfig{
		TableName:   "egg",
		IDField:     "EggID",
		RootIDField: "DuckID",
	}
}

func init() {
	db, err := NewMysqlDB(connectionString)
	if err != nil {
		panic(err)
	}

	std.SetTxProvider(db, nil)

	_, err = db.Exec(teardown)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(migrate)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(populate)
	if err != nil {
		panic(err)
	}

	repository, err = NewRepository(Duck{}, db)
	if err != nil {
		panic(err)
	}
}

func withProfile(next func(ctx context.Context) error) error {
	ctx := context.Background()
	profile := std.Profile{user.UserPermission{UserUUID: "123123123"}}
	return std.WithProfile(ctx, profile, func(ctx context.Context) error {
		return next(ctx)
	})
}

func TestMysqlRepository_Insert(t *testing.T) {
	err := withProfile(func(ctx context.Context) error {
		return std.WithRelTxContext(ctx, func(ctx context.Context) error {
			var d Duck

			err := gofakeit.Struct(&d)
			if err != nil {
				t.Error(err)
			}

			res, err := repository.Insert(ctx, &d)
			assert.NoError(t, err)
			assert.NotEqual(t, 0, res.ID)
			assert.NotEqual(t, "", res.UUID)

			return nil
		})
	})
	assert.NoError(t, err)
}

func TestMysqlRepository_GetByIDs(t *testing.T) {
	err := withProfile(func(ctx context.Context) error {
		ds := []Duck{}

		err := repository.GetByIDs(ctx, &ds, []int{3, 2, 4})
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 3, len(ds))
		assert.Equal(t, 3, *ds[0].DuckID)
		assert.Equal(t, 2, *ds[1].DuckID)
		assert.Equal(t, 4, *ds[2].DuckID)

		return nil
	})
	assert.NoError(t, err)
}

func TestMysqlRepository_FillStructByUUID(t *testing.T) {
	err := withProfile(func(ctx context.Context) error {
		var d Duck
		uuid := "0020c585-dfd2-3696-8245-6a211cb694a8"
		d.DuckUUID = &uuid

		err := repository.FillStructByUUID(ctx, &d)
		assert.NoError(t, err)

		assert.Equal(t, 3, len(d.Eggs))
		assert.Equal(t, "Lorine Nikolaus", *d.Eggs[0].Name)
		assert.Equal(t, "Garrison White", *d.Eggs[1].Name)
		assert.Equal(t, "Reta Skiles", *d.Eggs[2].Name)

		return nil
	})
	assert.NoError(t, err)
}

func TestMysqlRepository_Update(t *testing.T) {
	err := withProfile(func(ctx context.Context) error {
		uuid := "e5a4c708-81b6-3a98-9fcb-5ff72c5394ef"

		std.WithRelTxContext(ctx, func(ctx context.Context) error {
			var d Duck

			d.DuckUUID = &uuid
			d.Eggs = []Egg{}

			err := repository.Update(ctx, &d)
			assert.NoError(t, err)

			return nil
		})

		var newD Duck
		err := repository.GetByUUID(ctx, &newD, uuid)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(newD.Eggs))

		std.WithRelTxContext(ctx, func(ctx context.Context) error {
			var d Duck

			err := gofakeit.Struct(&d)
			if err != nil {
				panic(err)
			}

			d.DuckUUID = &uuid

			err = repository.Update(ctx, &d)
			assert.NoError(t, err)

			return nil
		})

		err = repository.GetByUUID(ctx, &newD, uuid)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(newD.Eggs))

		return nil
	})
	assert.NoError(t, err)
}

func TestMysqlRepository_Delete(t *testing.T) {
	err := withProfile(func(ctx context.Context) error {
		uuid := "1e1f0ffd-96ed-3c07-b73d-5d172ca2380b"

		err := std.WithRelTxContext(ctx, func(ctx context.Context) error {
			err := repository.Delete(ctx, uuid)
			assert.NoError(t, err)

			return nil
		})
		if err != nil {
			t.Error(err)
		}

		var d Duck
		err = repository.GetByUUID(ctx, &d, uuid)
		assert.Error(t, err)

		return nil
	})
	assert.NoError(t, err)
}
