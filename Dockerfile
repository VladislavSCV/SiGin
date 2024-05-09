# Use golang:alpine as a builder image
FROM golang:alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the application
RUN go build -o /app/app .


# Use a smaller image based on Alpine Linux
FROM alpine:latest

# Install ca-certificates package
RUN apk add --no-cache ca-certificates

# Set the working directory again
WORKDIR /app

# Copy the built binary from the builder image
COPY --from=build /app/app /app/app

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./app"]

