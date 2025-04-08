# Use Go base image
FROM golang:1.23

# Set working dir
WORKDIR /app

# Copy files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o main .

# Expose app port
EXPOSE 8080

# Run binary
CMD ["./main"]