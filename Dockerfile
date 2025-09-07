


FROM golang:1.24.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o otp-auth ./cmd/server

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/otp-auth .
EXPOSE 8080
CMD ["./otp-auth"]