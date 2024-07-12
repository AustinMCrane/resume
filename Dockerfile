FROM golang:1.21.3-alpine3.18 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go test -v ./...
RUN go build -o /app/bin/ .

FROM alpine:3.14
WORKDIR /app/bin
COPY --from=builder /app/bin/ /app/bin/
COPY --from=builder /app/assets /app/bin/assets
ENTRYPOINT ["./resume"]
