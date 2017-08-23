package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the version of s3-edit
const Version = "0.0.1"

// ShowVersion print the version of s3-edit
func ShowVersion() {
	fmt.Printf("s3-edit version: v%s\n", Version)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of s3-edit",
	Long: "Print the version of s3-edit",
	Run: func(cmd *cobra.Command, args []string) {
		ShowVersion()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
