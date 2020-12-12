# godeer
ディレクトリ構成を出力するcli

```
directory structure show and output csv

Usage:
  godeer [flags]
  godeer [command]

Available Commands:
  csv         dir structure output csv file. arg1: dirpath, arg2: savepath
  help        Help about any command
  show        show dir structure on array style. arg1: dirpath

Flags:
  -c, --char string   Select charcter code(utf-8, shift-jis) (default "utf-8")
  -h, --help          help for godeer
  -n, --nest int      Specify the depth of the directory (default 5)

Use "godeer [command] --help" for more information about a command.
```

## show
```
godeer show [ディレクトリパス]
```
指定したディレクトリパスの指定した階層の深さまでのファイル、ディレクトリをリスト表示します。
フラグによって表示する階層の深さ、文字コードの指定ができます。

## csv
```
godeer csv [ディレクトリパス] [csvファイル出力先]
```
指定したディレクトリパスの指定した階層の深さまでのファイル、ディレクトリを指定したcsvファイルとして出力する。
フラグによって表示する階層の深さ、文字コードの指定ができます。