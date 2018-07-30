FROM golang:1.10

# ENV GOPATH /go/src/app

RUN mkdir -p /go/src/app
WORKDIR /go/src/app/chi
COPY . /go/src/app

RUN apt-get update && \
    apt-get upgrade -y

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure

EXPOSE 8081

CMD go run server.go