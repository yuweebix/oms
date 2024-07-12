FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY . .

RUN rm Dockerfile docker-compose.yml Makefile example.txt README.md
RUN rm -fr bin/ db/ docs/ mocks/ tests/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go