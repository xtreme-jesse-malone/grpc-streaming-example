FROM golang:latest

ENV GO111MODULE=off
RUN go get github.com/xtreme-jesse-malone/grpc-streaming-example/server
RUN cd /go/src/github.com/xtreme-jesse-malone/grpc-streaming-example/server && go build .

ENTRYPOINT ["server"]
