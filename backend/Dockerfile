ARG VERSION=2.0.0
# Phase 1: Build Environment
FROM golang:1.22-alpine AS builder
WORKDIR /root/
RUN apk add --no-cache gcc musl-dev sqlite-dev
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o main .

# Phase 2: Runtime Environment
FROM alpine:3.20.3
WORKDIR /root/
EXPOSE 8080
COPY db /root/db
COPY .env /root/.env
COPY --from=builder /root/main .
CMD ["./main"]