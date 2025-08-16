# Use a builder image to compile your code
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]