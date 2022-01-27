package mdynamodb_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sun-fight/aws-client/mdynamodb"
	"github.com/sun-fight/aws-client/mdynamodb/pb"
)

func TestPutItem(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)
	var userInfo pb.UserInfo
	userInfo.UserID = 123
	userInfo.LastLoginAt = int32(time.Now().Unix())
	userInfo.CreatedAt = int32(time.Now().Unix())
	userInfo.Maps = map[string]*pb.Grid{
		"1": &pb.Grid{ID: 123, X: 1, Y: 2},
	}
	itemMap, err := attributevalue.MarshalMap(&userInfo)
	if err != nil {
		t.Fatal(err)
		return
	}
	exp, err := expression.NewBuilder().
		WithCondition(expression.AttributeNotExists(expression.Name("UserID"))).
		Build()
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(aws.ToString(exp.Condition()))
	fmt.Println(exp.Names())
	itemDao := mdynamodb.NewItemDao("User")
	_, err = itemDao.PutItem(mdynamodb.ReqPutItem{
		ItemMap:                  itemMap,
		ConditionExpression:      exp.Condition(),
		ExpressionAttributeNames: exp.Names(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetItem(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	itemDao := mdynamodb.NewItemDao("User")
	var userInfo pb.UserInfo
	_, err := itemDao.GetItem(mdynamodb.ReqGetItem{
		Key: map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
	}, &userInfo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(userInfo.String())
}

func TestUpdateItem(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	exp, err := expression.NewBuilder().
		WithUpdate(expression.Set(expression.Name("LastLoginAt"),
			expression.Value(time.Now().Unix()))).
		WithCondition(expression.AttributeExists(expression.Name("UserID"))).
		Build()
	if err != nil {
		t.Fatal(err)
		return
	}
	itemDao := mdynamodb.NewItemDao("User")
	_, err = itemDao.UpdateItem(mdynamodb.ReqUpdateItem{
		Key:                       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
		UpdateExpression:          exp.Update(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		ConditionExpression:       exp.Condition(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateItemByProto(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)

	updateItem := &pb.ReqTestUpdateItem{
		ExpUpdateOperations: []*pb.ExpUpdateOperation{
			&pb.ExpUpdateOperation{
				OperationMode: pb.EnumExpUpdateOperationMode_OperationModeSet,
				ExpUpdateSets: []*pb.ExpUpdateSet{
					&pb.ExpUpdateSet{
						Name: "LastLoginAt",
					},
				},
			},
		},
		UserInfo: &pb.UserInfo{
			LastLoginAt: int32(time.Now().Unix()),
		},
	}

	var updateBuilder expression.UpdateBuilder
	for _, v := range updateItem.ExpUpdateOperations {
		switch v.OperationMode {
		case pb.EnumExpUpdateOperationMode_OperationModeSet:
			for _, v2 := range v.ExpUpdateSets {
				// 获取 name,val
				name := expression.Name(v2.Name)
				var val expression.OperandBuilder
				switch v2.Name {
				case "LastLoginAt":
					val = expression.Value(updateItem.UserInfo.LastLoginAt)
				default:
				}
				switch v2.SetValMode {
				case pb.EnumExpUpdateSetValMode_SetValModePlus:
					val = name.Plus(val)
				case pb.EnumExpUpdateSetValMode_SetValModeIfNotExists:
					val = name.IfNotExists(val)
				default:
				}
				updateBuilder = updateBuilder.Set(name, val)
			}
		default:
		}
	}

	exp, err := expression.NewBuilder().WithUpdate(updateBuilder).
		WithCondition(expression.AttributeExists(expression.Name("UserID"))).
		Build()
	if err != nil {
		t.Fatal(err)
		return
	}
	itemDao := mdynamodb.NewItemDao("User")
	_, err = itemDao.UpdateItem(mdynamodb.ReqUpdateItem{
		Key:                       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
		UpdateExpression:          exp.Update(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		ConditionExpression:       exp.Condition(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteItem(t *testing.T) {
	initTestCfg()
	mdynamodb.Init(_cfg)
	itemDao := mdynamodb.NewItemDao("User")
	_, err := itemDao.DeleteItem(mdynamodb.ReqDeleteItem{
		Key: map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
	})
	if err != nil {
		t.Fatal(err)
	}
}
