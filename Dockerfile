FROM golang:latest
RUN mkdir /go/src/work
WORKDIR /go/src/work
CMD ["go","run","main.go"]
ADD . /go/src/work