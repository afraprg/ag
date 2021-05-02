FROM golang:1.16.3-buster

WORKDIR $GOPATH/src/github.com/afraprg/ag
RUN mkdir /root/config
ARG CONFIG_FILE="example_config.yml"
COPY $CONFIG_FILE /root/.ga.yaml
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .
ENTRYPOINT ["./ag"]