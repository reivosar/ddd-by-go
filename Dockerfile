FROM golang:latest as builder
COPY . /build/
WORKDIR /build

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

FROM alpine:3.15.4
WORKDIR /app
RUN adduser -S -D -H -h /app gouser
USER gouser

COPY --from=builder /build/main /app/

ENTRYPOINT ./main