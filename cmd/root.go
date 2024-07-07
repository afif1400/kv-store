package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "kv-store",
		Short: "A key-value store CLI application",
		Long:  `A key-value store CLI application that allows you to store and retrieve key-value pairs.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	godotenv.Load()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(jobCmd)
}
