FROM golang:1.21.3-alpine3.18 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/bin/ .

FROM alpine:3.14
COPY --from=builder /app/bin/ /app/bin/
ENTRYPOINT ["/app/bin/resume"]
