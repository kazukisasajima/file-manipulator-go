# file-manipulator-go

## 概要

引数を用いて簡単なファイル操作を行える goスクリプトです。ファイルの読み書きとシェルの引数入力の取り方について学習するために作成しました。このプロジェクトはコンピュータサイエンス学習サービス[Recursion](https://recursion.example.com)の課題でpythonで作成したものをGoで書き換えました。

## 内容

各コマンドの機能と使用方法を以下に示します。

| コマンド             | 使用方法                                                  |
| -------------------- | ------------------------------------------------------- |
| reverse            | go run file_manipulator.go reverse inputpath outputpath   |
| copy               | go run file_manipulator.go copy inputpath outputpath      |
| duplicate-contents | go run file_manipulator.go duplicate-contents inputpath n |
| replace-string     | go run file_manipulator.go replace-string inputpath a b   |

### 各コマンドの機能

- **`reverse`**: `inputpath` にあるファイルを受け取り、内容を逆にした新しいファイルを `outputpath` に作成します。
- **`copy`**: `inputpath` にあるファイルのコピーを作成し、`outputpath` として保存します。
- **`duplicate-contents`**: `inputpath` にあるファイルの内容を読み込み、その内容を複製し、複製された内容を `inputpath` に `n` 回複製します。
- **`replace-string`**: `inputpath` にあるファイルの内容から文字列 'a' を検索し、'a' の全てを 'b' に置き換えます。

## 使い方の例

### `reverse` コマンド

`input.txt` にあるファイルの内容を逆にして `output.txt` に保存します。

```sh
go run file_manipulator.go reverse input.txt output.txt
```
