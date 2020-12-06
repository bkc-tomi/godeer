package cmd

import (
	"fmt"
	"godeer/dirlist"
	"godeer/mymath"
	"strconv"

	"github.com/spf13/cobra"
)

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show dir structure on array style",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			/* 初期化 */
			var pathString string
			var nest int
			var err error
			var pathArray []dirlist.DirStruct
			var pathLen int

			/* 引数の取得 */
			pathString = args[0]
			nest, err = strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			/* パス配列 */
			// 取得
			tempArray, err := dirlist.GetDirArray(pathString, nest)

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
					tempDir,
					path.File,
				})
			}
			/* 画面出力 */
			for _, path := range pathArray {
				// ディレクトリの表示
				deep := 1
				for _, dir := range path.Dir {
					fmt.Printf("第%d階層: %10s, ", deep, dir)
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
