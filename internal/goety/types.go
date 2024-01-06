package goety

import "github.com/code-gorilla-au/goety/internal/logging"

type Service struct {
	logger logging.Logger
	dryRun bool
	client DynamoClient
	notify Notifier
}

type TableKeys struct {
	PartitionKey string
	SortKey      string
}
