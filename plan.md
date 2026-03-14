# このプログラムの概要
Goプロジェクトのテンプレートを作成するツール

# 使い方
gonew myproject
を実行すると
以下の構成でディレクトリとファイルを作成する

myproject/
  main.go <- main.go_ に書いてあるものが初期状態で書かれている
  go.mod <- go initしたときに作成されるやつ
  .gitignore <- .DS_Storeが書かれている
  .golangci.yml <- golangci-lintの設定ファイル
  Makefile <- Makefileのテンプレート

# 今後の予定
- [x] main.goのテンプレートを作成
- [x] go.modのテンプレートを作成
- [x] .gitignoreのテンプレートを作成
- [x] コマンドライン引数でプロジェクト名を受け取るようにする
- [x] ディレクトリとファイルを作成
- [x] すでに同名のディレクトリが存在する場合はエラーを出す
- [x] コマンドのビルドと実行の方法をドキュメントに書く