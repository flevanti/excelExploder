package main

import (
	"fmt"
	"github.com/flevanti/mongodbClient"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mdb, err := mongodbClient.Connect("mongodb+srv://appuser:appuser@cluster0-nnhfc.mongodb.net/test?retryWrites=true", "testdb");
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mdb.Close()
	collections, err := mdb.RetrieveCollectionsList()
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
	if len(collections) > 0 {
		fmt.Println("Collections found:")
		for _, collection := range collections {
			fmt.Println("- " + collection)
		}
	} else {
		fmt.Println("No collections found")
	}

	//fmt.Printf("hello %v", isAwsLambda.IsItInitialised())
	//fmt.Printf("hello %v", isAwsLambda.IsItLambda())
	//fmt.Printf("hello %v", isAwsLambda.IsItInitialised())

}
