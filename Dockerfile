#build stage
FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get -d -v ./...
RUN go build -o smart-fridge .

CMD ["/app/smart-fridge"]