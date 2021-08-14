FROM golang:1.16 as builder

# update
RUN apt-get update

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/simcart

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN GOOS=linux go build -a -o simcart .

# FROM alpine:latest

# RUN apk add ca-certificates

# RUN mkdir -p /bin/simcart
# WORKDIR /bin/simcart

# COPY --from=builder /go/src/github.com/simcart/.config.yml ./
# COPY --from=builder /go/src/github.com/simcart/run.sh ./
# COPY --from=builder /go/src/github.com/simcart/simcart ./

# RUN chmod +x ./simcart
CMD ["sh","./run.sh"]
