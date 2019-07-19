# Build binary
FROM golang:1.12-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git build-base
RUN go get -d -v ./...
RUN go test ./...
RUN go build -o bin/pricecalculator

# Run binary
FROM alpine:latest
ENV GIN_MODE release
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/bin/pricecalculator /pricecalculator
ENTRYPOINT [ "/pricecalculator" ]
EXPOSE 9096
