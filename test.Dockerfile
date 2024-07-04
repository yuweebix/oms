FROM golang:1.22-alpine
WORKDIR /app
RUN apk add --no-cache postgresql-client git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/vektra/mockery/v2@v2.43.2
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
