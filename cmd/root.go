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
		if isShowVersion {
			ShowVersion()
			return
		}
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

var isShowVersion bool
var awsProfile string

func init() {
	RootCmd.Flags().BoolVarP(&isShowVersion, "version", "v", false, "print the version of s3-edit")
	RootCmd.PersistentFlags().StringVarP(&awsProfile, "profile", "", "default", "Use a specific profile from your credential file")
}
