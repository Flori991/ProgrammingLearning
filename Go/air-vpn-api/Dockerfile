FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-server .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/go-server /go-server
CMD ["/go-server"]