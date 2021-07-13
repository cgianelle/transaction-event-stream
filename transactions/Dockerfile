FROM golang as builder
WORKDIR /usr/local/go/src/transactions
COPY go.mod .
COPY go.sum .
COPY transactions.go .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go install transactions


FROM alpine:latest
LABEL maintainer="cgianelle@gmail.com"
LABEL version="0.0.1"
COPY --from=builder /usr/local/go/bin/transactions /app/
WORKDIR /app

CMD ["./transactions"]