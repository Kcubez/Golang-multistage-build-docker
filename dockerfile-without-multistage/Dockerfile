# Dockerfile
FROM ubuntu as builder
# Install Go
RUN apt-get update && \
    apt-get install -y golang-go

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

ENV GO111MODULE=off

# Build the Go binary
RUN CGO_ENABLED=0 go build -o calculator calculator.go

# Run the binary
ENTRYPOINT ["./calculator"]