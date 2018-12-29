package main

import (
	"context"
	"github.com/flevanti/mongodbClient"
	"os"
)

var collRequestHeader = "request_header"
var collRequestRows = "request_rows"

func connectToDb() error {
	dbConn, err = mongodbClient.Connect(os.Getenv("DBCONNSTRING"), os.Getenv("DBDATABASE"))
	if err != nil {
		return err
	}

	//collections, err := dbConn.RetrieveCollectionsList()
	//if err != nil {
	//	return err
	//}
	//if len(collections) > 0 {
	//	fmt.Println("Collections found:")
	//	for _, collection := range collections {
	//		fmt.Println("- " + collection)
	//	}
	//} else {
	//	fmt.Println("No collections found")
	//}

	//result, err := dbConn.Db.Collection("hellothere").InsertMany(
	//	context.Background(),
	//	[]interface{}{bson.D{
	//		{"item", "canvas"},
	//		{"qty", 100},
	//		{"tags", bson.A{"cotton", "jeans"}},
	//		{"size", bson.D{
	//			{"h", 28},
	//			{"w", 35.5},
	//			{"uom", "cm"},
	//		}},
	//	},
	//		bson.D{
	//			{"item", "another canvas"},
	//			{"qty", 100},
	//			{"tags", bson.A{"plastic", "mercury"}},
	//			{"size", bson.D{
	//				{"h", 28},
	//				{"w", 35.5},
	//				{"uom", "cm"},
	//			}},
	//		},
	//	})
	//for _, i := range result.InsertedIDs {
	//	fmt.Println(i)
	//}

	return nil
}

func retrieveRequestInfoByRequestId(requestId string) string {
	var record interface{}
	//filter := bson.D{{"request_id", requestId}}
	var filter interface{}
	result := dbConn.Client.Database("testdb").Collection(collRequestHeader).FindOne(context.Background(), filter)
	err := result.Decode(record)
	if err != nil {
		return ""
	}
	return ""
}
