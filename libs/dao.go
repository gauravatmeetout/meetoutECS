package libs

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DAO struct {
	db               *mongo.Client
	connection_error error
	is_connected     bool
	Ctx              context.Context
	database         string
}

func ConnectToDB(username, password, hostman, database string) DAO {
	ctx := context.TODO()
	dao := DAO{nil, nil, false, ctx, database}
	dao.connect(username, password, hostman)
	return dao
}

func (dao *DAO) connect(username, password, hostname string) {
	password = url.QueryEscape(password)
	connetion_url := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", username, password, hostname)
	dao.db, dao.connection_error = mongo.NewClient(options.Client().ApplyURI(connetion_url))
	fmt.Println(connetion_url)
	if dao.connection_error != nil {
		fmt.Println(connetion_url)
		log.Fatal(dao.connection_error)
	}

	err := dao.db.Connect(dao.Ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (dao *DAO) GetDB() *mongo.Database {
	return dao.db.Database(dao.database)
}

func (dao DAO) ListDatabaseNames(to_close bool) {
	databases, err := dao.db.ListDatabaseNames(dao.Ctx, bson.M{})
	if err != nil {
		fmt.Println("Error in Listing")
		log.Fatal(err)
	}

	if to_close == true {
		defer dao.db.Disconnect(dao.Ctx)
	}
	fmt.Println(databases)
}
