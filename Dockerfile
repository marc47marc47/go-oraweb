FROM golang:1.4.2

COPY . /go/src/github.com/tapester/go-oraweb
WORKDIR /go/src/github.com/tapester/go-oraweb

RUN go get github.com/tools/godep

RUN godep restore
RUN godep go build && godep go install

EXPOSE 8081
CMD ["go-oraweb", "--bind", "0.0.0.0"]