# godeer
ディレクトリ構成を出力するcli

'''
$ godeer -help
Error: unknown shorthand flag: 'e' in -elp
Usage:
  godeer [flags]
  godeer [command]

Available Commands:
  csv         dir structure output csv file
  help        Help about any command
  show        show dir structure on array style

Flags:
  -h, --help   help for godeer

Use "godeer [command] --help" for more information about a command.

godeer: unknown shorthand flag: 'e' in -elp
'''

## show
'''
godeer show [ディレクトリパス] [階層]
'''
指定したディレクトリパスの指定した階層の深さまでのファイル、ディレクトリをリスト表示します。

## csv
'''
godeer csv [ディレクトリパス] [階層] [csvファイル出力先]
'''
指定したディレクトリパスの指定した階層の深さまでのファイル、ディレクトリを指定したcsvファイルとして出力する。