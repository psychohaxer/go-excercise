# Use specific version of Go base image
FROM golang:1.20

# Create working directory
WORKDIR /app

# Copy entire directory to working directory
COPY . /app

# Build Go application
RUN go build -o main .

# Specify default command to run when container is launched
CMD ["./main"]
