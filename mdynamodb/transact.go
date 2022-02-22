package mdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// transcatDao.ReqTransactGetItems = mdynamodb.ReqTransactGetItems{
// 	TransactItems: []types.TransactGetItem{
// 		{
// 			Get: &types.Get{
// 				Key:                  map[string]types.AttributeValue{"UserID": &types.AttributeValueMemberN{Value: "123"}},
// 				TableName:            aws.String("User"),
// 				ProjectionExpression: aws.String("UserID"),
// 			},
//		},
// 	},
// }
type ReqTransactGetItems struct {
	TransactItems []types.TransactGetItem
}

type ReqTransactWriteItems struct {
	TransactItems      []types.TransactWriteItem
	ClientRequestToken *string
}

type Transact struct {
}

func NewTransactDao() *Transact {
	return &Transact{}
}

func (item *Transact) TransactGetItems(req ReqTransactGetItems) (*dynamodb.TransactGetItemsOutput, error) {
	output, err := _client.TransactGetItems(context.TODO(), &dynamodb.TransactGetItemsInput{
		TransactItems: req.TransactItems,
	})
	if err != nil {
		return nil, err
	}
	if output.Responses == nil {
		err = ErrRecordNotFound
		return nil, err
	}
	return output, err
}

func (item *Transact) TransactWriteItems(req ReqTransactWriteItems) (*dynamodb.TransactWriteItemsOutput, error) {
	output, err := _client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: req.TransactItems,
	})
	if err != nil {
		return nil, err
	}
	return output, err
}
