package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/flevanti/isAwsLambda"
	"github.com/flevanti/mongodbClient"
	_ "github.com/joho/godotenv/autoload"
)

// this is the typical payload used by S3 to trigger aws lambda
// ExternalPayload* elements are added to use the same structure if we want to test it outside of AWS or use an external file

type PayloadType struct {
	RequestId string `json:"requestId"`
}

var dbConn *mongodbClient.ConnType
var err error

func main() {

	if isAwsLambda.IsItLambda() {
		// AWS lambda will add the payload to the handler call, we just need to specify the handler fn name...
		lambda.Start(Handler)
	} else {
		// because we are calling the handler manually, we load a payload (a dummy one)
		payload, err := LoadDummyPayload()
		if err != nil {
			fmt.Printf("Unable to load dummy payload: %s 💥 💥 💥 \n", err)
			return
		}
		output, err := Handler(payload)
		println(output)
		if err != nil {
			println("===============================")
			println("ERROR WHILE PROCESSING REQUEST:")
			println(err.Error())
			println("===============================")
		}
	}
}

func Handler(payload PayloadType) (string, error) {
	err = connectToDb()
	if err != nil {
		return "", err
	}
	defer dbConn.Close()

	var requestInfo string
	requestInfo = retrieveRequestInfoByRequestId(payload.RequestId)
	requestInfo = requestInfo

	return "OK!", nil
}

func moveExcelFileLocallyByRequestId(requestId string) (string, error) {
	return "", nil
}
