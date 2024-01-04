package dynamodb

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/code-gorilla-au/goety/internal/logging"
	"github.com/code-gorilla-au/odize"
)

func TestClient_Scan(t *testing.T) {
	logger := logging.New(false)
	ctx := logging.WithContext(context.Background(), logger)
	var client Client

	group := odize.NewGroup(t, nil)
	group.BeforeEach(func() {
		client = Client{
			logger: logger,
			db: &ddbClientMock{
				ScanFunc: func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
					return &dynamodb.ScanOutput{}, nil
				},
			},
		}
	})

	err := group.
		Test("should scan table", func(t *testing.T) {

			input := dynamodb.ScanInput{}
			_, err := client.Scan(ctx, &input)
			odize.AssertNoError(t, err)
		}).
		Run()

	odize.AssertNoError(t, err)
}
