package cmd

import (
	"fmt"
	"godeer/csv"
	"godeer/dirlist"
	"godeer/mymath"
	"strconv"

	"github.com/spf13/cobra"
)

func csvCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "csv",
		Short: "dir structure output csv file",
		Args:  cobra.RangeArgs(3, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			/* 初期化 */
			var err error
			var pathString string
			var nest int
			var savePath string
			var pathArray []dirlist.DirStruct
			var header []string
			var pathLen int

			/* 引数の取得 */
			pathString = args[0]

			nest, err = strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			savePath = args[2]

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
			// headerの作成
			for i := range pathArray[0].Dir {
				col := fmt.Sprintf("第%d階層", i+1)
				header = append(header, col)
			}
			header = append(header, "ファイル名")
			/* 出力 */
			err = csv.Write(savePath, header, pathArray)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
