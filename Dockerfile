FROM golang:1.16-alpine3.13

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app

ENTRYPOINT ["./app"]