FROM golang:1.16-buster AS build

WORKDIR /app

COPY .  .
COPY .env .env

ENV CGO_ENABLED=0
RUN go build -ldflags="-w -s" ./cmd/golang-apiserver


FROM scratch

WORKDIR /go/bin

COPY --from=build /app/golang-apiserver golang-apiserver
COPY --from=build /app/.env .env

EXPOSE 1234

ENTRYPOINT ["/go/bin/golang-apiserver"]
