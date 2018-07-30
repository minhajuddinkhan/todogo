FROM golang:1.9

RUN mkdir -p /app


COPY . /go/src/github.com/minhajuddinkhan/todogo

WORKDIR /go/src/github.com/minhajuddinkhan/todogo/cmd/todogo

RUN go get ./

EXPOSE 3000

CMD ["go", "run", "main.go"]

