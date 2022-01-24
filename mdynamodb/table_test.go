// Package mdynamodb_test provides
package mdynamodb_test

import (
	"aws-client/mdynamodb"
	"context"
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
	_cfg, err = config.LoadDefaultConfig(context.TODO(), // 本地的数据库 dynamodb
		config.WithRegion("us-east-2"),
		// dynamodb
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"97rwl", "oqzqpi", "")),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func TestDeleteTable(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)
	tableDao := mdynamodb.NewTableDao("CreditCardOffers")
	_, err := tableDao.DeleteTable()
	if err != nil {
		t.Fatal(err)
	}
}
