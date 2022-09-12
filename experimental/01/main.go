package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/infrastructure/util"
	"time"
)

func main() {

	const databaseName = "myjournal"

	uri := fmt.Sprintf("mongodb://localhost:27017/%s?readPreference=primary&ssl=false", databaseName)

	mwt := database.NewMongoWithTransaction(database.NewMongoDefault(uri), databaseName)

	mwt.PrepareCollection([]any{
		PurchaseOrder{},
	})

	testDB := TestDB{MongoWithTransaction: mwt}

	//testDB.Init(context.TODO())
	testDB.DoQuery(context.TODO())

}

type TestDB struct {
	*database.MongoWithTransaction
}

type PurchaseOrder struct {
	Product  string    `bson:"product" json:"product"`
	Total    int       `bson:"total" json:"total"`
	Customer string    `bson:"customer" json:"customer"`
	Date     time.Time `bson:"date" json:"date"`
}

func getTime(date string) time.Time {
	parse, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err.Error())
	}
	return parse
}

func (r *TestDB) Init(ctx context.Context) {

	_, err := r.SaveBulk(ctx, util.ToSliceAny([]PurchaseOrder{
		{Product: "Sikat Gigi", Total: 475, Customer: "Mirza", Date: getTime("2020-08-01")},
		{Product: "Gitar", Total: 19999, Customer: "Tika", Date: getTime("2020-08-02")},
		{Product: "Susu", Total: 1133, Customer: "Mirza", Date: getTime("2020-08-03")},
		{Product: "Pizza", Total: 850, Customer: "Omar", Date: getTime("2020-08-04")},
		{Product: "Sikat Gigi", Total: 475, Customer: "Omar", Date: getTime("2020-08-05")},
		{Product: "Pizza", Total: 475, Customer: "Zunan", Date: getTime("2020-08-06")},
		{Product: "Sikat Gigi", Total: 475, Customer: "Mirza", Date: getTime("2020-08-07")},
	}))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("init...\n")

}

func (r *TestDB) DoQuery(ctx context.Context) {

	coll := r.GetCollection(PurchaseOrder{})

	matchStage := bson.D{
		{
			"$match",
			bson.M{"customer": bson.M{"$in": bson.A{"Mirza"}}},
		},
	}

	groupStage := bson.D{
		{
			"$group",
			bson.M{
				"_id":   "$product",
				"total": bson.M{"$first": "$total"}},
		},
	}

	sortStage := bson.D{
		{
			"$sort",
			bson.M{"date": -1},
		},
	}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		groupStage,
		sortStage,
	})
	if err != nil {
		panic(err.Error())
	}

	for cursor.Next(ctx) {
		var result any
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf(">>>> %v\n", err.Error())
			panic(err.Error())
		}
		fmt.Printf(">>>> %v\n", util.MustJSON(result))
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf(">>>> %v\n", err.Error())
		panic(err.Error())
	}

}
