FROM golang:1.22-alpine
WORKDIR /app
RUN apk add --no-cache postgresql-client git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY . .
RUN find . -name '*_test.go' -delete \
  && find . -name '*_mock.go' -delete \
  && find . -type d -name mocks | xargs rm -rf
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go