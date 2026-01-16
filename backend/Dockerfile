FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build API binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend-api ./cmd/api

# Build worker binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend-worker ./cmd/worker

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/backend-api .
COPY --from=builder /app/backend-worker .

EXPOSE 8080

# Default to running API; docker-compose can override this
CMD ["./backend-api"]
