FROM golang:1.20-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app ./bin

EXPOSE 1234

CMD ["./app"]