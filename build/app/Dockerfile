FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./bin/quotes-service ./cmd/app/app.go

RUN chmod +x ./bin/quotes-service