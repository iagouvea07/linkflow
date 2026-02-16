FROM golang:1.24 AS builder

WORKDIR /app

COPY src/app/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o api

###
FROM golang:1.24 AS production

WORKDIR /app

COPY --from=builder /app/api /app/

ENTRYPOINT [ "/app/api" ]