# Use Go with Alpine Linux
FROM golang:1.23.3-alpine

# Install necessary dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy module files and download dependencies
COPY backend/go.mod backend/go.sum ./backend/
WORKDIR /app/backend
RUN go mod download

# Copy the rest of the project
COPY ./backend ./backend

# Build the Go binary
RUN go build -o EV-app ./cmd/main.go

# Expose port and set the command
EXPOSE 8000
CMD ["./EV-app"]
