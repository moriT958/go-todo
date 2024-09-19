# ステージ1
FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /app

# COPY go.sum
COPY go.mod main.go ./

RUN go mod download \
    && go build -o main /app/main.go

# ステージ2
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/main .

# 非特権ユーザーで実行
USER 1001

CMD [ "/app/main" ]
