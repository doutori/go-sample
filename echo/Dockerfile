FROM golang:1.10

# ENV GOPATH /go/src/app

RUN mkdir -p /go/src/app
WORKDIR /go/src/app/echo
COPY . /go/src/app

RUN apt-get update && \
    apt-get upgrade -y && \
    go get -u github.com/oxequa/realize && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure

EXPOSE 8080

CMD realize start --no-config --run