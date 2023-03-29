# Use a Go base image

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Install the Go dependencies
RUN go mod download
RUN apt-get update && apt-get install -y vim

# Copy the rest of the application files
COPY . .

# Build the Go application (main)

# Expose the right! port that the application will listen on
EXPOSE 8089  

# Run the application
CMD ["./main"]





