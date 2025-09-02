package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/code-gorilla-au/goety/internal/dynamodb"
	"github.com/code-gorilla-au/goety/internal/emitter"
	"github.com/code-gorilla-au/goety/internal/goety"
	"github.com/code-gorilla-au/goety/internal/logging"
	"github.com/code-gorilla-au/goety/internal/spinner"
	"github.com/spf13/cobra"
)

var (
	flagDumpTableName       string
	flagDumpEndpoint        string
	flagDumpFilePath        string
	flagDumpExtractAttrs    []string
	flagDumpLimit           int32
	flagDumpFilterExp       string
	flagDumpFilterAttrName  string
	flagDumpFilterAttrValue string
	flagDumpRawOutput       bool
)

var dumpCmd = &cobra.Command{
	Use:   "dump -t [TABLE_NAME] -p [FILE_PATH]",
	Short: "dump the contents of a dynamodb to a file",
	Long:  "dump will scan all items within a dynamodb table and write the contents to a file",
	Run:   dumpFunc,
}

func init() {
	dumpCmd.Flags().StringVarP(&flagDumpTableName, "table", "t", "", "table name")
	dumpCmd.Flags().StringVarP(&flagDumpEndpoint, "endpoint", "e", "", "DynamoDB endpoint to connect to, if none is provide it will use the default aws endpoint")
	dumpCmd.Flags().StringVarP(&flagDumpFilePath, "path", "p", "", "file path to save the json output")
	dumpCmd.Flags().StringSliceVarP(&flagDumpExtractAttrs, "attributes", "a", []string{}, "Optionally specify a list of attributes to extract from the table")
	dumpCmd.Flags().Int32VarP(&flagDumpLimit, "limit", "l", 0, "Limit the number of items returned per scan iteration")
	dumpCmd.Flags().StringVarP(&flagDumpFilterExp, "filter", "f", "", "Filter expression to apply to the scan operation")
	dumpCmd.Flags().StringVarP(&flagDumpFilterAttrName, "attribute-name", "N", "", "Filter expression attribute names")
	dumpCmd.Flags().StringVarP(&flagDumpFilterAttrValue, "attribute-value", "V", "", "Filter expression attribute values")
	dumpCmd.Flags().BoolVarP(&flagDumpRawOutput, "raw-output", "R", false, "Optional flag to output the dynamodb scan without transformation")
}

func dumpFunc(cmd *cobra.Command, args []string) {
	log := logging.New(flagRootVerbose)
	ctx := context.Background()

	if err := parseDumpFlag(); err != nil {
		log.Error("error parsing flags", "error", err)
		os.Exit(1)
	}

	log.Debug("loading dynamodb client")
	dbClient, err := dynamodb.NewClient(ctx, flagRootAwsRegion, flagDumpEndpoint)
	if err != nil {
		log.Error("could not load client")
		os.Exit(1)
	}

	msgEmitter := emitter.New()

	var writer goety.Writer
	if flagRootDryRun {
		log.Info("dry run enabled, no file will be created")
		writer = &bytes.Buffer{}
	} else {
		file, err := os.Create(flagDumpFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		writer = file
	}

	g := goety.New(dbClient, log, msgEmitter, flagRootDryRun)

	if !flagRootVerbose {
		spin := spinner.New(msgEmitter)
		spin.Start("starting dump")
		defer spin.Stop("dump complete")
	}
	_ = g.Dump(
		ctx,
		flagDumpTableName,
		writer,
		goety.WithAttrs(flagDumpExtractAttrs),
		goety.WithLimit(flagDumpLimit),
		goety.WithFilterExpression(flagDumpFilterExp),
		goety.WithFilterNameAttrs(flagDumpFilterAttrName),
		goety.WithFilterNameValues(flagDumpFilterAttrValue),
		goety.WithRawOutput(flagDumpRawOutput),
	)

}

// parsePurgeFlag will validate the flags passed to the purge command
func parseDumpFlag() error {
	if flagDumpTableName == "" {
		return errors.New("table name is required")
	}
	if flagDumpFilePath == "" {
		return errors.New("file path is required")
	}
	return nil
}
