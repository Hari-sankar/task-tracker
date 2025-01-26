FROM golang:alpine AS builder

WORKDIR /app

# Install required tools
RUN apk add --no-cache curl
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy and download dependencies
COPY go.mod ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Generate Swagger documentation
RUN swag init 


# Build the application
RUN go build -v -o main .

# Download and setup migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /app/migrate

# Final stage
FROM alpine:latest
WORKDIR /app

# Copy binaries and required files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY --from=builder /app/docs ./docs
COPY .env .
COPY migrate.sh .
COPY db/migrations ./migrations

# Set execute permissions
RUN chmod +x migrate.sh

EXPOSE 3000

ENTRYPOINT [ "/app/migrate.sh" ]
CMD ["/app/main","start"]
