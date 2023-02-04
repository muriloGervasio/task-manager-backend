FROM golang:1.20-alpine

RUN apk add --no-cache git

WORKDIR /app/task-manager

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/task-manager .

EXPOSE 8080

CMD ["./out/task-manager"]