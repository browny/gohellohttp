FROM golang:1.12.4-stretch as builder

WORKDIR /hello-http

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app 

FROM alpine:3.8

WORKDIR /root/

COPY --from=builder /hello-http/app .

CMD ["./app"]
