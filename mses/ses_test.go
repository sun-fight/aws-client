package mses_test

import (
	"aws-client/mses"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

var _cfg aws.Config

func initTestCfg() {
	// Load the Shared AWS Configuration (~/.aws/config)
	var err error
	_cfg, err = config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"97rwl", "oqzqpi", "")),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func TestDeleteTable(t *testing.T) {
	initTestCfg()
	mses.Init(_cfg, "from@email.com")
	newEmail := mses.NewEmailAll("subject", "test text", "<h1>test text</h1>", "to@email.com")
	res, err := newEmail.SendEmail()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
