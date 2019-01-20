FROM golang:1.11.0-stretch as builder

WORKDIR /

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:3.8

WORKDIR /root/

COPY --from=builder /main .

CMD ["./main"]