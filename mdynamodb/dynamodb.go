package mdynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	_client *dynamodb.Client
)

func Init(cfg aws.Config) {
	_client = dynamodb.NewFromConfig(cfg)
}

type dynamodbTable struct {
	tableName *string
}

// Key:                       map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
// UpdateExpression:          aws.String("SET Nickname = :u"),
// ExpressionAttributeValues: map[string]types.AttributeValue{":u": &types.AttributeValueMemberN{Value: "sun"}},
// ConditionExpression: "attribute_exists(UserID)",
type ReqUpdateItem struct {
	Key                       map[string]types.AttributeValue
	ConditionExpression       *string
	UpdateExpression          *string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
	ReturnValues              types.ReturnValue
}

// Key: map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
type ReqDeleteItem struct {
	Key                       map[string]types.AttributeValue
	ConditionExpression       *string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
	ReturnValues              types.ReturnValue
}

// ItemMap: map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
// ConditionExpression: "attribute_not_exists(UserID)",
type ReqPutItem struct {
	ItemMap                   map[string]types.AttributeValue
	ConditionExpression       *string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
}

// Key: map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
type ReqGetItem struct {
	Key            map[string]types.AttributeValue
	ConsistentRead *bool
}

// itemDao.ReqBatchGetItem = mdynamodb.ReqBatchGetItem{
// 	RequestItems: map[string]types.KeysAndAttributes{
// 		"User": {
// 			Keys: []map[string]types.AttributeValue{
// 				{
// 					"UserID": &types.AttributeValueMemberN{
// 						Value: "123",
// 					},
// 				},
// 			},
// 		},
// 	},
// }
type ReqBatchGetItem struct {
	RequestItems map[string]types.KeysAndAttributes
}

type ReqBatchWriteItem struct {
	RequestItems map[string][]types.WriteRequest
}

type dynamodbItem struct {
	tableName         *string
	ReqUpdateItem     ReqUpdateItem
	ReqDeleteItem     ReqDeleteItem
	ReqPutItem        ReqPutItem
	ReqGetItem        ReqGetItem
	ReqBatchGetItem   ReqBatchGetItem
	ReqBatchWriteItem ReqBatchWriteItem
}

func NewItemDao(tableName string) *dynamodbItem {
	return &dynamodbItem{
		tableName: aws.String(tableName),
	}
}

func NewTableDao(tableName string) *dynamodbTable {
	return &dynamodbTable{
		tableName: aws.String(tableName),
	}
}
