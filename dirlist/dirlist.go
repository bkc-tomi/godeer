package dirlist

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// LenError nest length error
type LenError struct {
	len int
}

func (e *LenError) Error() string {
	return fmt.Sprintf("正の値で指定してください。入力値: %d", e.len)
}

// GetDirArray 引数で指定されたパスのディレクトリ構造を配列で返す。
// 例）
// 引数：. 返り値：[dir dir_a][dir dir_b test.txt][dir dir_b test2.txt]
func GetDirArray(dir string, nest int) ([]DirStruct, error) {
	paths, err := dirwalk(dir, nest)
	if err != nil {
		return nil, err
	}
	pathArray := pathSeparator(paths)
	return pathArray, nil
}

// GetWinDirArray 引数で指定されたパスのディレクトリ構造を配列で返す。(文字コードはShift-JIS)
// 例）
// 引数：. 返り値：[dir dir_a][dir dir_b test.txt][dir dir_b test2.txt]
func GetWinDirArray(dir string, nest int) ([]DirStruct, error) {
	paths, err := dirwalk(dir, nest)
	if err != nil {
		return nil, err
	}
	// Shift-JISに変換
	sjPaths := utoSj(paths)
	pathArray := pathSeparator(sjPaths)

	return pathArray, nil
}

// dirwalk: パスで指定されたディレクトリ内の構造を配列として返す。
func dirwalk(dir string, nest int) ([]string, error) {
	/* 初期処理 */
	var paths []string

	/* 入力の取得 */
	// ディレクトリ情報の取得
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	/* 本処理 */
	if nest == 0 {
		// 指定の深さまで達したら現在までのパスを返す。
		paths = append(paths, dir)
		return paths, nil
	}

	// 指定の深さに達してなければさらに探索する。
	for _, file := range files {
		// ファイルパス文字列の作成
		filePath := filepath.Join(dir, file.Name())
		//
		if file.IsDir() && hasChild(filePath) {
			temp, _ := dirwalk(filePath, nest-1)
			paths = append(paths, temp...)
			continue
		}
		paths = append(paths, filePath)
	}
	/* 出力 */
	return paths, nil
}

// DirStruct ディレクトリ構造
type DirStruct struct {
	Dir  []string
	File string
}

// pathSeparator: パス文字列をセパレーターごとに分けて配列で返す。
// 例）dir/dir_b/test.txt　=> [dir dir_b test.txt]
func pathSeparator(paths []string) []DirStruct {
	/* 初期処理 */
	var sepPaths []DirStruct
	separator := string(filepath.Separator)

	/* 本処理 */

	for _, path := range paths {
		// ディレクトリとファイルを分離
		var dir, file string
		if filepath.Ext(path) != "" {
			// ファイルを指定するパスならディレクトリパスとファイル名に分ける
			dir, file = filepath.Split(path)
		} else {
			// ディレクトリを指定しているパスならfileは空文字
			dir = path
			file = ""
		}

		// DirStructに格納
		sepDir := strings.Split(dir, separator)
		sepPath := DirStruct{
			sepDir,
			file,
		}
		sepPaths = append(sepPaths, sepPath)
	}

	/* 出力 */
	return sepPaths
}

// utoSj utf-8 => shift-JISに文字コードを変換
func utoSj(strArray []string) []string {
	var results []string
	for _, str := range strArray {
		sjStr, _, _ := transform.String(japanese.ShiftJIS.NewEncoder(), str)
		results = append(results, sjStr)
	}
	return results
}

// hasChild 引数で指定されたディレクトリがディレクトリまたはファイルと言った子要素を持つかどうかを判定する
// 子要素を持つ場合はtrue, そうでない場合はfalseを返す。
func hasChild(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		return false
	}
	return true
}
