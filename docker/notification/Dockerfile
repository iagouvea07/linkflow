FROM golang:1.24 AS builder

WORKDIR /app

COPY src/notification/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

###
FROM golang:1.24 AS production

WORKDIR /app

COPY --from=builder /app/main /app/

ENTRYPOINT [ "/app/main" ]