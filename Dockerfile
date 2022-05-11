FROM golang:1.18.1-buster

WORKDIR $GOPATH/src/github.com/afraprg/ag
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .
ENTRYPOINT ["./ag"]