# サンプルプロジェクト

サンプルプロジェクトです。

## 機能

- xxx
- xxx
- xxx

## 前提条件

- Go 1.26以上

## Dockerで実行

DockerとDocker Compose、Task（またはMake）が必要です。

```bash
task build-linux
docker compose up --build
```

`make build-linux`でもビルドできます。Linux用バイナリは`CGO_ENABLED=0`で生成されます。
Dockerfileは`FROM scratch`を使用し、`bin/<プロジェクト名>-linux-amd64`をコピーして実行します。Composeの実行プラットフォームは`linux/amd64`です。
ARM64でネイティブ実行する場合は、Dockerfileのコピー元を`-linux-arm64`に、Composeの`platform`を`linux/arm64`に変更してください。
ソースを変更した場合は、Dockerイメージの作成前に`task build-linux`を再実行してください。

Buildx未導入の旧ビルダー（Classic Builder）でプラットフォーム不一致になる場合は、`docker compose config --images`でイメージ名を確認し、次の手順で起動してください。

```bash
docker build --platform linux/amd64 -t <確認したイメージ名> .
docker compose up --no-build
```

`scratch`にはシェルやCA証明書などが含まれません。HTTPS通信などが必要なアプリケーションでは、必要なファイルをDockerfileで追加してください。

```bash
docker compose down --rmi all --volumes --remove-orphans
```

docker compose で作成したリソースをすべて削除します。
