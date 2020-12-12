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
	RootCmd.PersistentFlags().IntP(
		"nest",
		"n",
		5,
		"Specify the depth of the directory",
	)
	RootCmd.PersistentFlags().StringP(
		"char",
		"c",
		"utf-8",
		"Select charcter code(utf-8, shift-jis)",
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
