package goety

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/code-gorilla-au/goety/internal/notify"
)

const (
	defaultBatchSize = 25
)

func New(client DynamoClient, logger *slog.Logger, notify Notifier, dryRun bool) Service {
	return Service{
		client: client,
		dryRun: dryRun,
		logger: logger,
		notify: notify,
	}
}

// Purge all items from the given table
//
// Example:
//
//	Purge(ctx, "my-table", TableKeys{ PartitionKey: "pk", SortKey: "sk" })
func (s Service) Purge(ctx context.Context, tableName string, keys TableKeys) error {
	s.logger.Debug("running purge")

	items, err := s.client.ScanAll(ctx, &dynamodb.ScanInput{
		TableName:       &tableName,
		AttributesToGet: []string{keys.PartitionKey, keys.SortKey},
	})
	if err != nil {
		s.logger.Error("could not scan table", "error", err)
		return err
	}

	if s.dryRun {
		s.logger.Debug("dry run enabled")
		prettyPrint(items)
		return nil
	}

	s.logger.Debug(fmt.Sprintf("purging %d items", len(items)))

	start := 0
	end := defaultBatchSize
	deleted := 0

	for start < len(items) {

		if end > len(items) {
			end = len(items)
		}

		batchItems := items[start:end]

		s.logger.Debug(fmt.Sprintf("batch delete %d items", len(batchItems)))
		_, err = s.client.BatchDeleteItems(ctx, tableName, batchItems)
		if err != nil {
			s.logger.Error("could not batch delete items", "error", err)
			return err
		}

		deleted += len(batchItems)
		start = end
		end += defaultBatchSize

		s.notify.Send(notify.Message{
			Message: fmt.Sprintf("deleted %d items", deleted),
		})

	}

	s.logger.Debug(fmt.Sprintf("purge complete, deleted: %d", deleted))

	return nil
}

// prettyPrint - prints a pretty json representation of the given value
func prettyPrint(v any) {
	data, err := json.MarshalIndent(v, "\n", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(data))
}
