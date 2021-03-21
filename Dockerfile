FROM golang:1.14

WORKDIR /go/src/app
COPY ./lab04.go .
RUN go get .
RUN go build
CMD ["app"]

