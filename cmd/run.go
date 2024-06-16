package cmd

import (
	"github.com/gkwa/heavybee/core"
	"github.com/spf13/cobra"
)

var jsonFile string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.RunBulkInsert(jsonFile)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&jsonFile, "json", "j", "", "JSON file containing data")
	if err := runCmd.MarkFlagRequired("json"); err != nil {
		panic(err)
	}
}
