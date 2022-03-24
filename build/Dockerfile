FROM golang:1.17-alpine

ARG SERVICE

WORKDIR /app

COPY . .

RUN go build -o main cmd/${SERVICE}/main.go

CMD ["./main"]