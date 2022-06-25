# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
RUN apk add linux-headers
RUN apk add build-base

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o /litra-web

EXPOSE 8080
CMD [ "/litra-web" ]