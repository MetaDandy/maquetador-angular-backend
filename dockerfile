FROM golang:1.24.1-alpine

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -tags netgo -ldflags '-s -w' -o app

CMD ["./app"]
