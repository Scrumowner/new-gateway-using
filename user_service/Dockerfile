FROM golang:1.21.5-alpine3.18 as builder

WORKDIR /usr/local/go/src/user_service

COPY . .


RUN go mod tidy
RUN go build -o main


FROM alpine:latest

COPY --from=builder /usr/local/go/src/user_service/main /main

EXPOSE 8081

RUN ["/main"]