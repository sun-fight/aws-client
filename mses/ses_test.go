package mses_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/sun-fight/aws-client/mses"
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

func TestSendEmail(t *testing.T) {
	initTestCfg()
	mses.Init(_cfg)
	newEmail := mses.NewEmail()
	res, err := newEmail.SendEmail(mses.NewReqSendEmail("subject", "test text", "<h1>test text</h1>", "to@email.com"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}
