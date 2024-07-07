# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.17 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Dependencies will be downloaded and installed to /go/pkg.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -o tui

# Use the official Debian-based image for a lean production container.
# https://hub.docker.com/_/debian
FROM debian:bullseye-slim

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/tui/tui

# Run the web service on container startup.
CMD ["/tui"]

