FROM golang:1.22.6

WORKDIR /app

COPY . .
RUN go mod download
RUN go build

ENTRYPOINT ["/app/go-mqtt-notifier"]