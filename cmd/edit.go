package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tsub/s3-edit/cli"
	"github.com/tsub/s3-edit/cli/s3"
)

var editCmd = &cobra.Command{
	Use:   "edit [S3 file path]",
	Short: "Edit directly a file on S3",
	Long: "Edit directly a file on S3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("required 1 argument")
			os.Exit(1)
		}

		path := s3.ParsePath(args[0])
		cli.Edit(path)
	},
}

func init() {
	RootCmd.AddCommand(editCmd)
}
