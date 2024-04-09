# Step 1: Build Go Application
FROM golang:1.21-alpine AS go-builder

WORKDIR /workspace

COPY . .

# Build the Go application
WORKDIR /workspace/apps/{{.serviceName}}

RUN go mod download
RUN go build -o app ./cmd

FROM alpine:latest

WORKDIR /app

# Copy Go application binary
COPY --from=go-builder /workspace/services/{{.serviceName}}/app .


# Expose port
EXPOSE 8080

# Set the entry point
CMD ["./app"]

