.PHONY: build build-linux build-windows run clean test install deps fmt lint

# バイナリ名
BINARY_NAME=gonew

# バージョン情報を埋め込むためのLDFLAGS
VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "")
ifeq ($(strip $(VERSION)),)
$(warning git describe --tags --abbrev=0 failed; VERSION will be empty)
endif

#GCFLAGS=-gcflags="-m"
GCFLAGS=
LDFLAGS=-ldflags "-s -w -X main.appVersion=$(VERSION)"
STRIP=-trimpath -buildvcs=false

# ビルドターゲット
build:
	@echo "Building..."
	go build $(GCFLAGS) $(STRIP) $(LDFLAGS) -o bin/$(BINARY_NAME) .

# Linux向けビルド
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build $(GCFLAGS) $(STRIP) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(GCFLAGS) $(STRIP) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 .
	@echo "Built binaries:"
	@ls -lh bin/$(BINARY_NAME)-linux-*

# Windows向けビルド
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build $(GCFLAGS) $(STRIP) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build $(GCFLAGS) $(STRIP) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-arm64.exe .
	@echo "Built binaries:"
	@ls -lh bin/$(BINARY_NAME)-windows-*

# 実行
run: build
	@echo "Running..."
	./bin/$(BINARY_NAME)

# クリーンアップ
clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

# テスト
test:
	@echo "Testing..."
	go test -v ./...

# インストール
install: build
	@echo "Installing to ~/bin/..."
	mkdir -p ~/bin
	cp bin/$(BINARY_NAME) ~/bin/
	@echo "Installed to ~/bin/$(BINARY_NAME)"

# 依存関係の更新
deps:
	@echo "Updating dependencies..."
	go mod tidy
	go mod download

# フォーマット
fmt:
	@echo "Formatting..."
	go fmt ./...

# リント
lint:
	@echo "Linting..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Installing..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...
