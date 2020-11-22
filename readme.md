# Account
家計簿ソフトです 
 
# DEMO 
"hoge"の魅力が直感的に伝えわるデモ動画や図解を載せる
 
# Features
以下の理由で動作が早いです
- golangを使用
- 並列処理
- テキストベース
 
# Requirement 
* go
 
# Installation 
[golangにてインストール](https://golang.org/)
 
# Usage
`data.txt`に日付・貸方科目・金額・借り方科目とスペースまたはタブで区切って入力する

コンパイルは、
```bash
go built account
```
実行は、
```
./account
```
# Note
貸借勘定科目を追加するには
53行目にお好きな貸借勘定科目を設定する
その際に科目数と同じ数の配列も設定する
ex.
```go
balance := [6]string{"現金","ゆうちょ普通","ゆうちょ定期","田舎銀行普通", "田舎銀行定期", "資本金"}
```
# Author
* 作成者 shibuya365
* shibuya365@gmail.com
 
# License
ライセンスを明示する
 
"Account" is under [MIT license](https://en.wikipedia.org/wiki/shibuya365).
 
"hoge" is not Confidential.