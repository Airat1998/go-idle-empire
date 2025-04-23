FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-idle-empire ./cmd

EXPOSE 8080

CMD ["./go-idle-empire"]
