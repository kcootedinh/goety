package commands

import "github.com/spf13/cobra"

var (
	flagRootVerbose   = false
	flagRootDryRun    = false
	flagRootAwsRegion = "ap-southeast-2"
)

var rootCmd = &cobra.Command{
	Use:   "goety [COMMAND] --[FLAGS]",
	Short: "dynamodb purge tool",
	Long:  "dynamodb purge tool",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagRootVerbose, "verbose", "v", false, "add verbose logging")
	rootCmd.PersistentFlags().BoolVarP(&flagRootDryRun, "dry-run", "d", false, "dry run mode outputs items to sdt out")
	rootCmd.PersistentFlags().StringVarP(&flagRootAwsRegion, "aws-region", "r", "ap-southeast-2", "aws region, default is ap-southeast-2")

	rootCmd.AddCommand(purgeCmd)
}

func Execute() error {

	return rootCmd.Execute()
}
