package db

import (
	"context"

	"github.com/qiniu/qmgo"
)

func NewMongoConnection() *qmgo.Collection {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri: "mongodb://localhost:27017",
	})

	db := client.Database("technical-test-pt-semesta-arus-technology")
	coll := db.Collection("scrapping")

	if err := client.Ping(10); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	return coll
}
