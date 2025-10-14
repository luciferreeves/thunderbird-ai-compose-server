FROM golang:1.25 AS builder

WORKDIR /thunderbird-ai-compose-server

RUN apt-get update && apt-get install -y gcc libc6-dev make git

ENV CGO_ENABLED=1

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o thunderbird-ai-compose-server

FROM debian:bookworm-slim

WORKDIR /thunderbird-ai-compose-server

RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

COPY --from=builder /thunderbird-ai-compose-server/thunderbird-ai-compose-server .

CMD ["./thunderbird-ai-compose-server"]
