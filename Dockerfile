# Use an official Go runtime as a base image
FROM golang:1.21.0

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files into the container
COPY go.mod go.sum ./

# Install dependencies using Go modules (if applicable)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o myapp .

# Set environment variables
ENV PORT=8000

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./myapp"]