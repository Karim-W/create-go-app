# Step 1: Build Go Application
FROM golang:1.21-alpine AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd

FROM alpine:latest

WORKDIR /app

# Copy Go application binary
COPY --from=go-builder /app/app .


# Expose port
EXPOSE 8080

# Set the entry point
CMD ["./app"]
