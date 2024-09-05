# Build
FROM golang:1.21-alpine AS builder

WORKDIR /src

COPY . .

RUN go build -o baojia .

# Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/baojia .

EXPOSE 8080

ENTRYPOINT ["./baojia"]