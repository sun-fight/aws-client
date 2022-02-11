package mdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (item *dynamodbItem) GetItem(req ReqGetItem, out interface{}) (*dynamodb.GetItemOutput, error) {
	output, err := _client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:            req.Key,
		ConsistentRead: req.ConsistentRead,
		TableName:      item.tableName,
	})
	if err != nil {
		return nil, err
	}
	if output.Item == nil {
		err = ErrGetItemNotFound
		return nil, err
	}
	err = attributevalue.UnmarshalMap(output.Item, &out)
	return output, err
}

// key := map[string]types.AttributeValue{
// 	"UserID" :&types.AttributeValueMemberN{
// 		Value: "11211",
// 	},
// }
func (item *dynamodbItem) PutItem(req ReqPutItem) (output *dynamodb.PutItemOutput, err error) {
	return _client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:                 item.tableName,
		Item:                      req.ItemMap,
		ConditionExpression:       req.ConditionExpression,
		ExpressionAttributeNames:  req.ExpressionAttributeNames,
		ExpressionAttributeValues: req.ExpressionAttributeValues,
	})
}

func (item *dynamodbItem) UpdateItem(req ReqUpdateItem) (output *dynamodb.UpdateItemOutput, err error) {
	return _client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 item.tableName,
		Key:                       req.Key,
		ConditionExpression:       req.ConditionExpression,
		UpdateExpression:          req.UpdateExpression,
		ExpressionAttributeNames:  req.ExpressionAttributeNames,
		ExpressionAttributeValues: req.ExpressionAttributeValues,
		ReturnValues:              req.ReturnValues,
	})
}

func (item *dynamodbItem) DeleteItem(req ReqDeleteItem) (output *dynamodb.DeleteItemOutput, err error) {
	return _client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName:                 item.tableName,
		Key:                       req.Key,
		ConditionExpression:       req.ConditionExpression,
		ExpressionAttributeNames:  req.ExpressionAttributeNames,
		ExpressionAttributeValues: req.ExpressionAttributeValues,
		ReturnValues:              req.ReturnValues,
	})
}

func (item *dynamodbItem) Query(req ReqQueryInput) (output *dynamodb.QueryOutput, err error) {
	return _client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 item.tableName,
		IndexName:                 req.IndexName,
		KeyConditionExpression:    req.KeyConditionExpression,
		FilterExpression:          req.FilterExpression,
		ExpressionAttributeNames:  req.ExpressionAttributeNames,
		ExpressionAttributeValues: req.ExpressionAttributeValues,
		ExclusiveStartKey:         req.ExclusiveStartKey,
		Limit:                     req.Limit,
	})
}

func (item *dynamodbItem) BatchGetItem(req ReqBatchGetItem) (output *dynamodb.BatchGetItemOutput, err error) {
	return _client.BatchGetItem(context.TODO(), &dynamodb.BatchGetItemInput{
		RequestItems: req.RequestItems,
	})
}

func (item *dynamodbItem) BatchWriteItem(req ReqBatchWriteItem) (output *dynamodb.BatchWriteItemOutput, err error) {
	return _client.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: req.RequestItems,
	})
}
