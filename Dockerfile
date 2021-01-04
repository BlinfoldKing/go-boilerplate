FROM golang:1.15

ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH

RUN go get -u golang.org/x/lint/golint

