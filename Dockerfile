# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the host to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port on which the app will run
EXPOSE 8080

# Set the command to run the executable
CMD ["./main"]
