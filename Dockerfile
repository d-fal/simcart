FROM golang:1.16

# update
RUN apt-get update

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/simcart

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /simcart 



CMD ["run.sh"]
