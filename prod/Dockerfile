FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc musl-dev

COPY ["./storyteller-api/go.mod", "./storyteller-api/go.sum", "./"]
RUN go mod download

COPY storyteller-api ./
COPY .env .env
RUN go build -o bin/storyteller-api cmd/api/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/storyteller-api /

COPY .env .env

CMD ["/storyteller-api"]
