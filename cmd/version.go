package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kv-store",
	Long:  `All software has versions. This is kv-store's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`kv-store v0.1 -- HEAD`)
	},
}
