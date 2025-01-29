FROM golang:1.23

WORKDIR /app

COPY  . .

RUN apt-get update && apt-get install -y librdkafka-dev
RUN go build -o /app/cmd/walletcore/goapp /app/cmd/walletcore/main.go

ENTRYPOINT ["/app/cmd/walletcore/goapp"]
# CMD ["tail", "-f", "/dev/null"]