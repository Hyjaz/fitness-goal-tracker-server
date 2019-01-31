FROM golang:1.11.0-stretch as builder
WORKDIR /
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM scratch
WORKDIR /
ENV DBHOSTNAME="mongodb-service" DBPORTNUMBER="27017" DBNAME="fitness-goal-tracker"
COPY --from=builder /main .
CMD ["./main"]