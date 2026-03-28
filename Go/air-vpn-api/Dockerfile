FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-server .

FROM scratch
COPY --from=builder /app/go-server /go-server
EXPOSE 3000
CMD ["/go-server"]