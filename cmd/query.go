package cmd

import (
	"github.com/gkwa/heavybee/core"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query items from DynamoDB",
	Long:  `This command queries all items from the DynamoDB table.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.QueryItems()
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
