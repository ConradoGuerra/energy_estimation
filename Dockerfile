FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine
WORKDIR /app
COPY --from=builder /app/ .
EXPOSE 8080
CMD ["./main"]