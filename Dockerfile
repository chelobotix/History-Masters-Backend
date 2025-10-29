
FROM golang:1.25-alpine

RUN apk add --no-cache git curl bash build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" \
    github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o go-history-masters ./cmd/api

EXPOSE 4000 2345

# overwrite by docker compose command:
CMD ["air", "-c", ".air.toml"]