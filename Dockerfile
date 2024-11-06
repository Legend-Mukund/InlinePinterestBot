FROM golang:1.23.2-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/inline cmd/main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/bin/inline ./bin/inline

CMD ["./bin/inline"]