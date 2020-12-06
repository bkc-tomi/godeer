package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		showCmd(),
		csvCmd(),
	)
}

// RootCmd ルートコマンド
var RootCmd = &cobra.Command{
	Use:   "godeer",
	Short: "directory structure show and output csv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("godeer")
	},
}
