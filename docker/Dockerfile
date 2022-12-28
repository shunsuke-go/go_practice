FROM golang:1.19.4-alpine3.17
RUN apk update && apk add git && apk add make
ENV GO111MODULE=on
RUN mkdir /go/src/go_practice
COPY . /go/src/go_practice
WORKDIR /go/src/go_practice/src
RUN go install github.com/cosmtrek/air@v1.40.4
RUN go build -o ./main.go

CMD ["air", "-c", ".air.toml"]