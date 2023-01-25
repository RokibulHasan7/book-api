# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest AS builder
#FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Md. Rokibul Hasan <mdrokibulhasan18@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

## Copy go mod and sum files
#COPY go.mod go.sum ./
#
## Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container


# Build the Go app
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 go build -a -o main

#CMD ["./main"]

########### Final Stage ############
FROM alpine:latest

WORKDIR /root/

COPY --from=builder  /app/main .

# Expose port 3333 to the outside world
EXPOSE 3333

# Command to run the executable
CMD ["./main"]