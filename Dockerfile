# ステージ1: ビルド用
FROM golang:1.23-alpine AS build

# ワーキングディレクトリを設定
WORKDIR /app

# モジュールファイルをコピー
COPY go.mod go.sum ./
RUN go mod tidy

# ソースコードを全てコピー
COPY . .

# Goアプリケーションをビルド
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
