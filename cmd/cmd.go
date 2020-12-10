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
	RootCmd.PersistentFlags().IntP("nest", "n", 10, "表示するディレクトリの深さを指定")
	RootCmd.PersistentFlags().StringP("os", "o", "win", "csv出力の際のosごとの文字エンコード")
}

// RootCmd ルートコマンド
var RootCmd = &cobra.Command{
	Use:   "godeer",
	Short: "directory structure show and output csv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("godeer")
	},
}
