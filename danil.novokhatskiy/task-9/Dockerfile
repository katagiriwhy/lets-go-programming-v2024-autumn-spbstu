FROM golang:1.16.3-alpine3.13

WORKDIR /task-9

COPY . .

RUN go get -d -v ./...

RUN go build -o api ./cmd/service/main.go

EXPOSE 8080

CMD ["./api"]