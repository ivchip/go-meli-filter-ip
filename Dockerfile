FROM golang:1.14.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Build the Go app
RUN go build -o go-meli-filter-ip .

# Run the executable
CMD ./go-meli-filter-ip