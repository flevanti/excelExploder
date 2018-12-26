package main

import (
	"fmt"
	"github.com/flevanti/isAwsLambda"
	_ "github.com/joho/godotenv/autoload"
	"mongoClient"
)

func main() {
	connected, err := mongodbClient.Connect("mongodb:fdfdsfdfs", "fdsfs");
	if err != nil || connected != true {
		fmt.Println(err)
		return
	}

	fmt.Printf("hello %v", isAwsLambda.IsItInitialised())
	fmt.Printf("hello %v", isAwsLambda.IsItLambda())
	fmt.Printf("hello %v", isAwsLambda.IsItInitialised())
}
