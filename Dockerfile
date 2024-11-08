# Stage 1: Build the server
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o shortify-server

# Stage 2: Serve
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/shortify-server .
EXPOSE 5566
CMD ["./shortify-server"]
