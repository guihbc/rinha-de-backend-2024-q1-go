FROM golang:1.21.7-alpine3.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/api ./cmd/main.go

FROM busybox:1.36.1

COPY --from=builder /app/bin/api .
CMD ["/api"]
