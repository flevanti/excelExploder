package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/flevanti/goUtils"
	"github.com/flevanti/isAwsLambda"
	"github.com/flevanti/mongodbClient"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

// this is the typical payload used by S3 to trigger aws lambda
// LocalPayload* elements are added to use the same structure if we want to test it outside of AWS
type PayloadType struct {
	LocalPayload     bool   `json:"localPayload"`
	LocalPayloadFile string `json:"localPayloadFile"`
	Records          []struct {
		EventVersion string    `json:"eventVersion"`
		EventSource  string    `json:"eventSource"`
		AwsRegion    string    `json:"awsRegion"`
		EventTime    time.Time `json:"eventTime"`
		EventName    string    `json:"eventName"`
		UserIdentity struct {
			PrincipalID string `json:"principalId"`
		} `json:"userIdentity"`
		RequestParameters struct {
			SourceIPAddress string `json:"sourceIPAddress"`
		} `json:"requestParameters"`
		ResponseElements struct {
			XAmzRequestID string `json:"x-amz-request-id"`
			XAmzID2       string `json:"x-amz-id-2"`
		} `json:"responseElements"`
		S3 struct {
			S3SchemaVersion string `json:"s3SchemaVersion"`
			ConfigurationID string `json:"configurationId"`
			Bucket          struct {
				Name          string `json:"name"`
				OwnerIdentity struct {
					PrincipalID string `json:"principalId"`
				} `json:"ownerIdentity"`
				Arn string `json:"arn"`
			} `json:"bucket"`
			Object struct {
				Key       string `json:"key"`
				Size      int    `json:"size"`
				ETag      string `json:"eTag"`
				VersionID string `json:"versionId"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

var dbConn mongodbClient.ConnType
var err error

func main() {

	if isAwsLambda.IsItLambda() {
		// AWS lambda will add the payload to the handler call, we just need to specify the handler fn name...
		// If the lambda is the local docker implementation, we need to pass the payload (a dummy one) as an argument
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
			println("============================")
			println("ERROR WHILE PROCESSING REQUEST:")
			println(err.Error())
			println("============================")
		}
	}
}

func Handler(payload PayloadType) (string, error) {
	err = connectToDb()
	if err != nil {
		return "", err
	}

	return "OK!", nil
}

func LoadDummyPayload() (PayloadType, error) {
	dummyPayloadFile := os.Getenv("DUMMYPAYLOADFILE")
	if len(dummyPayloadFile) == 0 {
		return PayloadType{}, errors.New("dummy payload filename not found in configuration file")
	}
	if goUtils.FileExists(dummyPayloadFile) == false {
		return PayloadType{}, errors.New(fmt.Sprintf("dummy payload filename provided `%s` not found", dummyPayloadFile))
	}
	dummypayloadfileContent, err := goUtils.ReadFileContent(dummyPayloadFile)
	if err != nil {
		return PayloadType{}, err
	}
	var payload PayloadType
	err = json.Unmarshal([]byte(dummypayloadfileContent), &payload)
	if err != nil {
		return PayloadType{}, err
	}
	return payload, nil
}

func connectToDb() error {
	dbConn, err := mongodbClient.Connect(os.Getenv("DBCONNSTRING"), os.Getenv("DBDATABASE"))
	if err != nil {
		return err
	}
	defer dbConn.Close()

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
