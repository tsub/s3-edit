package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd define root command
var RootCmd = &cobra.Command{
	Use:   "s3-edit",
	Short: "Edit directly a file on Amazon S3",
	Long:  "Edit directly a file on Amazon S3",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute execute root command
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
