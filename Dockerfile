# Go 1.23のAlpineイメージを使用してビルドステージを作成
FROM golang:1.23-alpine AS build

# Dockerコンテナ内の作業ディレクトリを /app に設定
WORKDIR /app

# ローカルの go.mod と go.sum ファイルをコンテナ内の /app にコピー
COPY go.mod go.sum ./

# コンテナ内で go mod tidy を実行し、依存関係を整理・インストール     
RUN go mod tidy

# ローカルのソースコード全てをコンテナ内の /app にコピー
COPY . .

# コンテナ内の /app/cmd/app ディレクトリでビルドし、main という実行ファイルを作成
RUN go build -o main ./cmd/app

# ステージ2: 実行用
FROM alpine:latest

# 証明書ファイルを含むディレクトリを作成（GoのHTTPS通信に必要）
RUN apk --no-cache add ca-certificates

# ビルドされたバイナリをコピー
COPY --from=build /app/main /app/main

# 環境変数の読み込みに使用される.envをコピー
COPY .env /app/.env


# アプリケーションを起動
CMD ["/app/main"]
