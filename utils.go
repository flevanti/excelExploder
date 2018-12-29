package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flevanti/goUtils"
	"os"
)

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
