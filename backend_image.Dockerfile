# Start from base image
FROM golang:alpine

# Set the current working directory inside the container
WORKDIR /backend

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source from current directory to working directory
COPY . .

# Build the application
RUN go build -o mvc ./cmd/main.go


# Copy the wait-for.sh script
COPY wait-for.sh .

# Give execute permission to wait-for.sh
RUN chmod +x wait-for.sh

# Expose necessary port
EXPOSE 8000

# Run the created binary executable after wait for mysql container to be up
CMD ["./wait-for.sh" , "mysql:3306" , "--timeout=300" , "--" , "./mvc"]