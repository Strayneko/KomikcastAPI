FROM golang:alpine

# set workdir
WORKDIR /usr/src/app

# copy go mod and go sum
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# copy app source code
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o server .

# starting server
ENTRYPOINT ["./server"]