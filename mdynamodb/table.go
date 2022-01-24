package mdynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (item *dynamodbTable) ListTables() {
	// Build the input parameters for the request.
	input := &dynamodb.ListTablesInput{
		Limit: aws.Int32(10),
	}

	resp, err := _client.ListTables(context.TODO(), input)
	if err != nil {
		panic("failed to list table, " + err.Error())
	}

	for _, v := range resp.TableNames {
		fmt.Println("table " + v)
	}
}

// DeleteTable 删除表
func (item *dynamodbTable) DeleteTable() (*dynamodb.DeleteTableOutput, error) {
	output, err := _client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: item.tableName,
	})
	return output, err
}
