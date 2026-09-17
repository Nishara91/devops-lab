# Stage 1: Build the Go binary
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o server main.go

# Stage 2: Create a minimal lightweight production image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY index.html .

EXPOSE 8080
CMD ["./server"]