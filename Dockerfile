# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for some modules)
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs 

EXPOSE 8083

CMD ["./main"]
