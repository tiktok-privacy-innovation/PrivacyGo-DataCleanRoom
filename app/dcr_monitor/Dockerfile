FROM golang:1.22 as builder

WORKDIR /app

COPY ./github.com /go/pkg/mod/github.com
COPY ./client ./client
COPY ./monitor ./monitor
COPY ./go.mod ./go.mod
COPY ./main.go ./main.go
COPY ./go.sum ./go.sum
RUN go build -o /app/dcr_monitor && chmod +x /app/dcr_monitor

FROM debian:12.4
RUN useradd dcrmonitor
RUN apt-get update && apt-get install -y ca-certificates
COPY ./conf /usr/local/dcr_conf

WORKDIR /app
RUN chown -R dcrmonitor:dcrmonitor /app

COPY --from=builder --chown=dcrmonitor:dcrmonitor /app/dcr_monitor /usr/local/bin/dcr_monitor
COPY ./conf /usr/local/dcr_conf
USER dcrmonitor 
ENTRYPOINT dcr_monitor