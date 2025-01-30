FROM golang:1.23

WORKDIR /app

COPY  . .

RUN apt-get update && apt-get install -y librdkafka-dev
RUN go build -o /app/balancesapp/cmd/goapp /app/balancesapp/cmd/main.go

ENTRYPOINT ["/app/balancesapp/cmd/goapp"]
# CMD ["tail", "-f", "/dev/null"]