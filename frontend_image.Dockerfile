# Use an official Go runtime as the base image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o mvc ./cmd/main.go

# Use a minimal lightweight image to serve the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install bash
RUN apk add --no-cache bash

# Copy the built Go application from the build stage to the current directory
COPY --from=build /app/mvc .
RUN chmod +x mvc

# Copy your HTML templates, CSS, and JS files to the container
COPY templates/ ./templates/
COPY templates/static/ ./static/ 

# Expose the port that your application runs on
EXPOSE 8080

# Define the command to run your application when the container starts
CMD ["bash", "-c", "ls -l && pwd && if [ -f ./mvc ]; then echo 'Executable exists'; ./mvc; else echo 'Executable not found'; fi"]
