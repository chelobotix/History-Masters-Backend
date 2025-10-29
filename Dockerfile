
FROM golang:1.25-alpine

RUN apk add --no-cache git curl bash build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN go build -o go-history-masters ./cmd/api

EXPOSE 4000 2345

CMD ["air", "-c", ".air.toml"]