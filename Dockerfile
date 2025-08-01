FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
