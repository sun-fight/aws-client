package mdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (item *dynamodbItem) GetItem(out interface{}) (*dynamodb.GetItemOutput, error) {
	output, err := _client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:            item.ReqGetItem.Key,
		ConsistentRead: item.ReqGetItem.ConsistentRead,
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
func (item *dynamodbItem) PutItem() (output *dynamodb.PutItemOutput, err error) {
	return _client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:                 item.tableName,
		Item:                      item.ReqPutItem.ItemMap,
		ConditionExpression:       item.ReqPutItem.ConditionExpression,
		ExpressionAttributeNames:  item.ReqPutItem.ExpressionAttributeNames,
		ExpressionAttributeValues: item.ReqPutItem.ExpressionAttributeValues,
	})
}

func (item *dynamodbItem) UpdateItem() (output *dynamodb.UpdateItemOutput, err error) {
	return _client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 item.tableName,
		Key:                       item.ReqUpdateItem.Key,
		ConditionExpression:       item.ReqUpdateItem.ConditionExpression,
		UpdateExpression:          item.ReqUpdateItem.UpdateExpression,
		ExpressionAttributeNames:  item.ReqUpdateItem.ExpressionAttributeNames,
		ExpressionAttributeValues: item.ReqUpdateItem.ExpressionAttributeValues,
		ReturnValues:              item.ReqUpdateItem.ReturnValues,
	})
}

func (item *dynamodbItem) DeleteItem() (output *dynamodb.DeleteItemOutput, err error) {
	return _client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName:                 item.tableName,
		Key:                       item.ReqDeleteItem.Key,
		ConditionExpression:       item.ReqDeleteItem.ConditionExpression,
		ExpressionAttributeNames:  item.ReqDeleteItem.ExpressionAttributeNames,
		ExpressionAttributeValues: item.ReqDeleteItem.ExpressionAttributeValues,
		ReturnValues:              item.ReqDeleteItem.ReturnValues,
	})
}

func (item *dynamodbItem) BatchGetItem() (output *dynamodb.BatchGetItemOutput, err error) {
	return _client.BatchGetItem(context.TODO(), &dynamodb.BatchGetItemInput{
		RequestItems: item.ReqBatchGetItem.RequestItems,
	})
}

func (item *dynamodbItem) BatchWriteItem() (output *dynamodb.BatchWriteItemOutput, err error) {
	return _client.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: item.ReqBatchWriteItem.RequestItems,
	})
}
