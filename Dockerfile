FROM hub.hamdocker.ir/golang:alpine

WORKDIR /app

COPY . /app

RUN go install

CMD ["/go/bin/ping-service"]