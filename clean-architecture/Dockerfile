FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/clean-architecture-service cmd/api/main.go

FROM alpine:3.18

RUN apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/clean-architecture-service .

COPY --from=builder /app/config.json .
EXPOSE 8080

CMD ["./clean-architecture-service"]