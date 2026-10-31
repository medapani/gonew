# gonew

Goプロジェクトのテンプレートを作成するCLIツールです。

## ビルド

```bash
# Makefile
make build

# Taskfile
task build
```

## インストール

```bash
# Makefile
make install

# Taskfile
task install
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
- `Taskfile.yml`
- `Dockerfile`
- `docker-compose.yml`
- `README.md`
- `internal/`
- `cmd/`

## 生成したプロジェクトのDocker実行

```bash
cd myproject
task build-linux
docker compose up --build
```

`make build-linux`でもLinux用バイナリを生成できます。Dockerfileは`FROM scratch`を使用し、`bin/myproject-linux-amd64`を`/myproject`へコピーして起動します。Composeの実行プラットフォームも`linux/amd64`に設定します。

Docker内ではGoのビルドを行わないため、ソースを変更した場合は先に`task build-linux`を再実行してください。

Buildx未導入の旧ビルダー（Classic Builder）でプラットフォーム不一致になる場合は、次のように明示的にイメージを作成してから起動してください。

```bash
docker build --platform linux/amd64 -t myproject-app .
docker compose up --no-build
```

`myproject-app`はComposeのデフォルトのイメージ名です。ディレクトリ名やComposeのプロジェクト名を変更した場合は、`docker compose config --images`で確認してください。

```bash
docker compose down --rmi all --volumes --remove-orphans
```

docker compose で作成したリソースをすべて削除します。
