FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download && go mod verify

COPY . .

RUN [ "go", "run", "main.go" ]