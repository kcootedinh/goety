package dynamodb

import (
	"context"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

//go:generate moq -rm -stub -out mocks_test.go . ddbClient
type ddbClient interface {
	Scan(ctx context.Context, params *ddb.ScanInput, optFns ...func(*ddb.Options)) (*ddb.ScanOutput, error)
	BatchWriteItem(ctx context.Context, params *ddb.BatchWriteItemInput, optFns ...func(*ddb.Options)) (*ddb.BatchWriteItemOutput, error)
}
