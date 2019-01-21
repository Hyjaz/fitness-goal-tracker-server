FROM golang:1.11.0-stretch as builder

WORKDIR /

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM scratch

WORKDIR /root/

ENV DBHOSTNAME "db"

COPY --from=builder /main .

CMD ["./main"]