FROM golang:1.15-alpine

LABEL maintainer="Parithiban G <parithiban1@gmail.com>"

WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./main

CMD ["./main"]