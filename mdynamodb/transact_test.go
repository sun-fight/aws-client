// Package mdynamodb_test provides
package mdynamodb_test

import (
	"aws-client/mdynamodb"
	"aws-client/mdynamodb/model"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestTransactGetItems(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	transcatDao := mdynamodb.NewTransactDao()
	transcatDao.ReqTransactGetItems = mdynamodb.ReqTransactGetItems{
		TransactItems: []types.TransactGetItem{
			{Get: &types.Get{
				Key:                  map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				TableName:            aws.String("User"),
				ProjectionExpression: aws.String("UserID"),
			}},
		},
	}
	out, err := transcatDao.TransactGetItems()
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range out.Responses {
		var user model.User
		err = attributevalue.UnmarshalMap(v.Item, &user)
		fmt.Println(k)
		fmt.Println(err)
		fmt.Println(user)
	}
}

func TestTransactWriteItems(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	transcatDao := mdynamodb.NewTransactDao()
	transcatDao.ReqTransactWriteItems = mdynamodb.ReqTransactWriteItems{
		TransactItems: []types.TransactWriteItem{
			{
				ConditionCheck: &types.ConditionCheck{
					TableName: aws.String("User"),
					Key:       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				},
				Put: &types.Put{
					TableName: aws.String("User"),
					Item:      map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				},
				Delete: &types.Delete{
					TableName: aws.String("User"),
					Key:       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				},
				Update: &types.Update{
					TableName: aws.String("User"),
					Key:       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				},
			},
		},
	}
	_, err := transcatDao.TransactWriteItems()
	if err != nil {
		t.Fatal(err)
	}
}
