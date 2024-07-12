# Dockerfile.multi

# Stage 1: Build the Go binary
FROM ubuntu as builder

# Install Go
RUN apt-get update && \
    apt-get install -y golang-go

# Set working directory
WORKDIR /app

# Set environment variables
ENV GO111MODULE=off

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 go build -o calculator calculator.go

# Stage 2: Create the final image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/calculator /calculator

# Run the binary
ENTRYPOINT ["/calculator"]

