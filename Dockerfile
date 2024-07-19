FROM golang:1.22-bookworm

ENV GOOSE_MIGRATION_DIR=/app/db/migrations
ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING="postgres://username:password@transfer-db:5432/transfer-db?sslmode=disable"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/air-verse/air@latest

RUN go build -o /app/main ./cmd/main.go

EXPOSE 8081

CMD ["air", "-c", "/app/.air.toml"]
