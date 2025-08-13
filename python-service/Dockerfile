# Stage 1: Build the Go app
FROM golang:1.22-alpine3.18 

WORKDIR /app


# Copy all source files
COPY . .

# Build the Go binary
RUN go build -o python-service .

RUN apk add python3


# Run the Go binary
CMD ["./python-service","-i"]
