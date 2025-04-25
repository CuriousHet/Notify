# Use official Golang image
FROM golang:1.23

# Create app directory
WORKDIR /app

# Copy everything
COPY . .

# Download Go modules
RUN go mod download

# Build the app
RUN go build -o notify .

# Expose ports for gRPC (5050) and HTTP (GraphQL + Prometheus - 8081)
EXPOSE 5050
EXPOSE 8081

# Run the app
CMD ["./notify"]
