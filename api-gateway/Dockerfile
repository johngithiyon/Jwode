FROM golang:1.22-alpine3.18 

WORKDIR /app


# Copy all source files
COPY . .

# Build the Go binary
RUN go build -o api



# Run the Go binary
CMD ["./api"]
