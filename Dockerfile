FROM golang:1.16-buster AS build

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod ./
COPY go.sum ./
COPY .  .
COPY .env .env
RUN go mod download
RUN go mod tidy
RUN go mod verify

RUN go build -ldflags="-w -s" ./cmd/golang-apiserver


FROM scratch

COPY --from=build /app/golang-apiserver /go/bin/golang-apiserver
COPY --from=build /app/.env /go/bin/.env

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/golang-apiserver"]
