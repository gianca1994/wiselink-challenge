FROM golang:1.18

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 8080
CMD ["./wiselink"]