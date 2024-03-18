FROM golang:1.22.1-alpine

WORKDIR /app

COPY . .

RUN go build -o interpreter .

ENTRYPOINT ["/app/interpreter"]


