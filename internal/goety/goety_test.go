package goety

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/code-gorilla-au/goety/internal/logging"
	"github.com/code-gorilla-au/odize"
)

func TestService_Purge(t *testing.T) {
	var client DynamoClientMock
	var service Service
	logger := logging.New(true)
	ctx := logging.WithContext(context.Background(), logger)

	callScanAll := 0
	callBatchDelete := 0

	group := odize.NewGroup(t, nil)

	group.BeforeEach(func() {

		client = DynamoClientMock{
			ScanAllFunc: func(ctx context.Context, input *dynamodb.ScanInput) ([]map[string]types.AttributeValue, error) {
				callScanAll++
				return []map[string]types.AttributeValue{
					{
						"pk": &types.AttributeValueMemberS{Value: "pk"},
						"sk": &types.AttributeValueMemberS{Value: "sk"},
					},
				}, nil
			},
			BatchDeleteItemsFunc: func(ctx context.Context, tableName string, keys []map[string]types.AttributeValue) (*dynamodb.BatchWriteItemOutput, error) {
				callBatchDelete++
				return &dynamodb.BatchWriteItemOutput{}, nil
			},
		}

		service = Service{
			client: &client,
			dryRun: false,
			logger: logger,
		}
	})

	group.AfterEach(func() {
		callScanAll = 0
		callBatchDelete = 0
	})

	err := group.
		Test("should purge items", func(t *testing.T) {
			err := service.Purge(ctx, "my-table", TableKeys{PartitionKey: "pk", SortKey: "sk"})
			odize.AssertNoError(t, err)

			odize.AssertEqual(t, 1, callScanAll)
			odize.AssertEqual(t, 1, callBatchDelete)
		}).
		Test("should not delete items on dry run", func(t *testing.T) {
			service.dryRun = true

			err := service.Purge(ctx, "my-table", TableKeys{PartitionKey: "pk", SortKey: "sk"})
			odize.AssertNoError(t, err)

			odize.AssertEqual(t, 1, callScanAll)
			odize.AssertEqual(t, 0, callBatchDelete)
		}).
		Test("should return error if scan fails", func(t *testing.T) {
			expectedErr := errors.New("scan all error")
			client.ScanAllFunc = func(ctx context.Context, input *dynamodb.ScanInput) ([]map[string]types.AttributeValue, error) {
				return nil, expectedErr
			}

			err := service.Purge(ctx, "my-table", TableKeys{PartitionKey: "pk", SortKey: "sk"})
			odize.AssertTrue(t, errors.Is(err, expectedErr))
		}).
		Test("should return error if batch write fails", func(t *testing.T) {
			expectedErr := errors.New("batch write error")
			client.BatchDeleteItemsFunc = func(ctx context.Context, tableName string, keys []map[string]types.AttributeValue) (*dynamodb.BatchWriteItemOutput, error) {
				return nil, expectedErr
			}

			err := service.Purge(ctx, "my-table", TableKeys{PartitionKey: "pk", SortKey: "sk"})
			odize.AssertTrue(t, errors.Is(err, expectedErr))
		}).
		Run()

	odize.AssertNoError(t, err)

}
