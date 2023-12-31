FROM golang:1.21-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main cmd/app/main.go

EXPOSE 8080

CMD ["/app/main"]
