# Start with the base Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules file and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source code to the container
COPY . .

# Build the Go program inside the container
RUN go build -o myprogram

# Set the entrypoint for the container
ENTRYPOINT ["./myprogram"]
