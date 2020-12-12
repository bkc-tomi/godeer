package cmd

import (
	"fmt"
	"godeer/dirlist"
	"godeer/mymath"

	"github.com/spf13/cobra"
)

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show dir structure on array style. arg1: dirpath",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			/* 初期化 */
			var pathString string
			var nest int
			var flg string
			var err error
			var pathArray []dirlist.DirStruct
			var pathLen int

			/* 引数の取得 */
			pathString = args[0]

			nest, err = cmd.Flags().GetInt("nest")
			if err != nil {
				return err
			}

			flg, err = cmd.Flags().GetString("char")
			if err != nil {
				return err
			}

			/* パス配列 */
			// 取得
			tempArray, err := dirlist.GetDirArray(pathString, nest, flg)

			if err != nil {
				fmt.Println(err)
				return nil
			}
			// 整形
			for _, path := range tempArray {
				// 一番深いディレクトリの深さを取得
				pathLen = mymath.IntMax(pathLen, len(path.Dir))
			}
			for _, path := range tempArray {
				dl := len(path.Dir)
				tempDir := path.Dir
				// 深いディレクトリに合わせて空文字を配列に追加
				if dl < pathLen {
					for i := 0; i < pathLen-dl; i++ {
						tempDir = append(tempDir, "")
					}
				}
				// 整形したディレクトリ構造をpathArrayに渡す
				pathArray = append(pathArray, dirlist.DirStruct{
					Dir:  tempDir,
					File: path.File,
				})
			}
			/* 画面出力 */
			for _, path := range pathArray {
				//　文字コードの変換
				var tempDir []string
				switch flg {
				case "utf-8":
					tempDir = path.Dir
				case "shift-jis":
					tempDir = dirlist.UtoSj(path.Dir)
				default:
					tempDir = path.Dir
				}

				deep := 1
				// ディレクトリの表示
				for _, dir := range tempDir {
					fmt.Printf("第%d階層: %25s, ", deep, dir)
					deep++
				}
				// ファイルの表示
				fmt.Printf("ファイル: %s \n", path.File)
			}
			return nil
		},
	}

	return cmd
}
