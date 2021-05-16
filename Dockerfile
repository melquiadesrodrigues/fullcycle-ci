FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build math.go

CMD ["./math"]