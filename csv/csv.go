package csv

import (
	"encoding/csv"
	"godeer/dirlist"
	"os"
)

// Write pathに指定したファイルにdataを書き込んで出力
func Write(path string, header []string, data []dirlist.DirStruct) error {
	/* 初期処理 */
	var err error
	var file *os.File
	/* ファイルを開く */
	file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	/* 処理終了後ファイルを閉じる */
	defer file.Close()

	/* ファイルを空にする */
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	/* データの書き込み */
	// ライターの作成
	writer := csv.NewWriter(file)
	// ライターバッファ
	err = writer.Write(header)
	if err != nil {
		return err
	}
	for _, dirpath := range data {
		pathArray := dirpath.Dir
		pathArray = append(pathArray, dirpath.File)
		err = writer.Write(pathArray)
		if err != nil {
			return err
		}
	}
	// 書き込み
	writer.Flush()

	return nil
}
