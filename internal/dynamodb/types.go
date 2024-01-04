package dynamodb

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Client - dynamodb client to query the table (get,put,query,scan)
type Client struct {
	db     *dynamodb.Client
	logger *slog.Logger
	dryRun bool
}
