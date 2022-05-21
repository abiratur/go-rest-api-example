FROM golang:1.18-alpine

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY passwords.txt ./

RUN go build -o ./app

EXPOSE 8080

ENTRYPOINT ["./app"]