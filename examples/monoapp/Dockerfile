# Step 1: Build Go Application
FROM golang:1.22-alpine AS go-builder

WORKDIR /workspace

COPY . .

# Build the Go application
WORKDIR /workspace/apps/{{.serviceName}}

RUN go mod download
RUN go build -o app ./cmd

FROM alpine:latest

WORKDIR /app

# Copy Go application binary
COPY --from=go-builder /workspace/apps/{{.serviceName}}/app .
COPY --from=go-builder /workspace/apps/{{.serviceName}}/internal/sql_migrations ./internal/sql_migrations
COPY --from=go-builder /workspace/apps/{{.serviceName}}/swagger ./swagger


# Expose port
EXPOSE 8080

# Set the entry point
CMD ["./app"]
