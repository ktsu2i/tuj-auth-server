FROM golang:1.23-bookworm

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/rubenv/sql-migrate/...@latest

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app
COPY .air.toml /app/.air.toml

CMD ["sh", "-c", "sql-migrate up && air -c .air.toml"]
