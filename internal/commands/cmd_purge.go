package commands

import (
	"context"
	"os"

	"github.com/code-gorilla-au/goety/internal/dynamodb"
	"github.com/code-gorilla-au/goety/internal/goety"
	"github.com/code-gorilla-au/goety/internal/logging"
	"github.com/spf13/cobra"
)

var (
	flagPurgeTableName    string
	flagPurgeEndpoint     string
	flagPurgePartitionKey string
	flagPurgeSortKey      string
)

var purgeCmd = &cobra.Command{
	Use:   "purge -t [TABLE_NAME] -p [PARTITION_KEY] -s [SORT_KEY]",
	Short: "purge a dynamodb table of all items",
	Long:  "purge a dynamodb table of all items",
	Run:   purgeFunc,
}

func init() {
	purgeCmd.PersistentFlags().StringVarP(&flagPurgeTableName, "table", "t", "", "table name")
	purgeCmd.PersistentFlags().StringVarP(&flagPurgeEndpoint, "endpoint", "e", "", "DynamoDB endpoint to connect to, if none is provide it will use the default aws endpoint")
	purgeCmd.PersistentFlags().StringVarP(&flagPurgePartitionKey, "partition-key", "p", "pk", "The name of the partition key, default is pk")
	purgeCmd.PersistentFlags().StringVarP(&flagPurgeSortKey, "sort-key", "s", "sk", "The name of the sort key, default is sk")
}

// purgeFunc is the entry point for the purge command. It will purge a dynamodb table of all items
func purgeFunc(cmd *cobra.Command, args []string) {
	log := logging.New(flagRootVerbose)
	ctx := context.Background()

	log.Debug("loading dynamodb client")
	dbClient, err := dynamodb.NewClient(ctx, flagRootAwsRegion, flagPurgeEndpoint)
	if err != nil {
		log.Error("could not load client")
		os.Exit(1)
	}

	g := goety.New(dbClient, log, flagRootDryRun)

	if err = g.Purge(ctx, flagPurgeTableName, goety.TableKeys{
		PartitionKey: flagPurgePartitionKey,
		SortKey:      flagPurgeSortKey,
	}); err != nil {
		log.Error("error purging table", "error", err)
		os.Exit(1)
	}

}
