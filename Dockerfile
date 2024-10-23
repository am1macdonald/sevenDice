# Development stage
FROM golang:1.23-alpine AS dev

# Install Air for hot reloading and build essentials
RUN go install github.com/air-verse/air@latest && \
    apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Command to run Air
CMD ["air"]
