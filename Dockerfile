# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o server ./main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Add certificates for HTTPS
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server .
#COPY --from=builder /app/config.yaml .

EXPOSE 8080

CMD ["./server"]
