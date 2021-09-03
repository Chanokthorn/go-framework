package std

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	mysqlTxKey = "mysqlTx"
	mongoTxKey = "mongoTx"
)

var (
	mysql       *sqlx.DB
	mongoClient *mongo.Client
)

// SetTxProvider sets global mysql db, mongo client, etc. for later usage. Must be set on application starting up
// for example:
//	db, err := mysql.NewDB(config.MySQL)
//	if err != nil {
//		log.Fatalf("Error on opening database connection: %s", err.Error())
//	}
//	mongoClient, mongoDB := mongo.NewClientAndDB(config.Mongo.Uri, config.Mongo.DatabaseName)
//	std.SetTxProvider(db.DB, mongoClient.Client)
func SetTxProvider(db *sqlx.DB, client *mongo.Client) {
	mysql = db
	mongoClient = client
}

// WithRelDocTxContext Used as closure function. Will provide ctx with Mysql and Mongo context that must be retrieved
// using hooks such as UseMysqlTx and UseMongoTx, this only works if SetTxProvider has already been invoked
// for example:
//	err = std.WithRelDocTxContext(ctx, func(ctx std.Context) error {
//		productID, err = handler.productService.CreateTx(ctx, prd)
//		if err != nil {
//			return err
//		}
//		return nil
//	})
func WithRelDocTxContext(ctx context.Context, next func(ctx context.Context) error) error {
	mySqlTx := mysql.MustBegin()

	s, err := mongoClient.StartSession()
	if err != nil {
		return nil
	}

	sc := mongo.NewSessionContext(ctx, s)

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	if err = s.StartTransaction(txnOpts); err != nil {
		return err
	}

	ctx = context.WithValue(ctx, mysqlTxKey, mySqlTx)

	ctx = context.WithValue(ctx, mongoTxKey, &sc)

	err = next(ctx)
	if err != nil {
		if errRollBack := mySqlTx.Rollback(); errRollBack != nil {
			return errRollBack
		}
		if errRollBack := rollbackMongoTx(ctx, sc); errRollBack != nil {
			return errRollBack
		}
		return err
	}

	err = mySqlTx.Commit()
	if err != nil {
		return err
	}

	err = commitMongoTx(ctx, sc)
	if err != nil {
		return err
	}

	return nil
}

// WithRelTxContext Used as closure function. Will provide ctx with Mysql context that must be retrieved
// using UseMysqlTx hook, this only works if SetTxProvider has already been invoked
func WithRelTxContext(ctx context.Context, next func(ctx context.Context) error) error {
	mySqlTx := mysql.MustBegin()

	ctx = context.WithValue(ctx, mysqlTxKey, mySqlTx)

	err := next(ctx)
	if err != nil {
		if errRollBack := mySqlTx.Rollback(); errRollBack != nil {
			return errRollBack
		}
	}

	err = mySqlTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// WithRelDocTxContext Used as closure function. Will provide ctx with Mongo context that must be retrieved
// using UseMongoTx hook, this only works if SetTxProvider has already been invoked
func WithDocTxContext(ctx context.Context, next func(ctx context.Context) error) error {
	s, err := mongoClient.StartSession()
	if err != nil {
		return nil
	}

	sc := mongo.NewSessionContext(ctx, s)

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	if err = s.StartTransaction(txnOpts); err != nil {
		return err
	}

	ctx = context.WithValue(ctx, mongoTxKey, &sc)

	err = next(ctx)
	if err != nil {
		if errRollBack := rollbackMongoTx(ctx, sc); errRollBack != nil {
			return errRollBack
		}
		return err
	}

	err = commitMongoTx(ctx, sc)
	if err != nil {
		return err
	}

	return nil
}

// UseMysqlTx Retrieves *sqlx.Tx from ctx created from  WithRelDocTxContext
// for example:
//	sqlxTx, err := std.UseMysqlTx(ctx)
//	if err != nil {
//		return fmt.Errorf(`unable to get tx: %v`, err)
//	}
func UseMysqlTx(ctx context.Context) (*sqlx.Tx, error) {
	mysqlTx, ok := ctx.Value(mysqlTxKey).(*sqlx.Tx)
	if !ok {
		return nil, fmt.Errorf(`unable to retrieve mysqlTx from context`)
	}

	return mysqlTx, nil
}

// UseMongoTx Retrieves *mongo.SessionContext from ctx created from  WithRelDocTxContext
// for example:
//	sessionContext, err := std.UseMongoTx(ctx)
//	if err != nil {
//		return fmt.Errorf(`unable to cast ctx to session context: %v`, err)
//	}
func UseMongoTx(ctx context.Context) (mongo.SessionContext, error) {
	mongoTx, ok := ctx.Value(mongoTxKey).(*mongo.SessionContext)
	if !ok {
		return nil, fmt.Errorf(`unable to retrieve mongoTx from context`)
	}

	return *mongoTx, nil
}

func commitMongoTx(ctx context.Context, tx mongo.SessionContext) error {
	defer tx.EndSession(ctx)

	if err := tx.CommitTransaction(ctx); err != nil {
		return fmt.Errorf(`unable to commit mongo transaction: %v`, err)
	}

	return nil
}

func rollbackMongoTx(ctx context.Context, tx mongo.SessionContext) error {
	defer tx.EndSession(ctx)

	if err := tx.AbortTransaction(ctx); err != nil {
		return fmt.Errorf(`unable to abort mongo transaction: %v`, err)
	}

	return nil
}
