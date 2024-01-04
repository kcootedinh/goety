package dynamodb

import (
	"errors"
	"log/slog"
)

var (
	ErrNoItems = errors.New("no items found")
)

// Client - dynamodb client to query the table (get,put,query,scan)
type Client struct {
	db     ddbClient
	logger *slog.Logger
	dryRun bool
}
