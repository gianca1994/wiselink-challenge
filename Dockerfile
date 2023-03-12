FROM golang:1.19-alpine

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

ENV API_PORT=8080 \
    DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=root \
    DB_NAME=wiselink \
    DB_SSL_MODE=disable \
    JWT_SECRET=secret \
    JWT_TIME_EXPIRE_MINUTES=60

CMD ["./app"]
