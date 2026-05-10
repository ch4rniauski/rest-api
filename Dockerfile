FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o api ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api .
COPY migrations/ ./migrations/

EXPOSE 8080

CMD ["/app/api"]
