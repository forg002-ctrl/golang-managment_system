FROM golang:1.21.5-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o golang_api ./cmd/app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/golang_api .
COPY --from=builder /app/.env .
EXPOSE 8200
CMD ["./golang_api"]
