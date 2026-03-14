# gonew

Goプロジェクトのテンプレートを作成するCLIツールです。

## ビルド

```bash
go build -o gonew .
```

## 実行

```bash
./gonew myproject
```

## バージョン表示

```bash
./gonew -v
```

これを実行すると、`myproject/` 配下に以下が作成されます。

- `main.go` (`main.go_` のテンプレート)
- `go.mod` (`go mod init` で生成)
- `.git/` (`git init` で生成)
- `.gitignore` (`.DS_Store`)
- `.golangci.yml`
- `Makefile`
- `README.md`
- `internal/`
- `cmd/`
