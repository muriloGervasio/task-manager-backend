FROM golang:1.20-alpine

RUN apk add --no-cache git

WORKDIR /app/task-manager

ENV PORT=$port

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/task-manager .

EXPOSE $PORT

CMD ["./out/task-manager"]