# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Diana Lucia Serna Higuita <dserna@gmx.net.com>"

# Set the Current Working Directory inside the container
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app


# Copy the source from the current directory to the Working Directory inside the container
COPY service/ /usr/src/app

RUN chown -R 777 /usr/src/app

# Build the Go app
RUN go build -o main .


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Delete the source code
RUN rm -r pkg/ go.mod go.sum main.go

# Expose port 4060 to the outside world
EXPOSE 4060

# Command to run the executable
CMD ["./main"]
