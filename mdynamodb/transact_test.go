// Package mdynamodb_test provides
package mdynamodb_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sun-fight/aws-client/mdynamodb"
	"github.com/sun-fight/aws-client/mdynamodb/pb"
)

func TestTransactGetItems(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	dao := mdynamodb.NewTransactDao()
	out, err := dao.TransactGetItems(mdynamodb.ReqTransactGetItems{
		TransactItems: []types.TransactGetItem{
			{Get: &types.Get{
				Key:                  map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
				TableName:            aws.String("User"),
				ProjectionExpression: aws.String("UserID"),
			}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range out.Responses {
		var userInfo pb.UserInfo
		err = attributevalue.UnmarshalMap(v.Item, &userInfo)
		fmt.Println(k)
		fmt.Println(err)
		fmt.Println(userInfo.String())
	}
}

func TestTransactWriteItems(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	dao := mdynamodb.NewTransactDao()
	_, err := dao.TransactWriteItems(mdynamodb.ReqTransactWriteItems{
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
	})
	if err != nil {
		t.Fatal(err)
	}
}
