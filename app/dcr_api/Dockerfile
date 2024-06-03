FROM golang:1.22 as builder
WORKDIR /app
COPY ./biz ./biz
COPY ./main.go ./main.go
COPY ./router_gen.go ./router_gen.go
COPY ./router.go ./router.go
# TODO DELETE
COPY ./github.com /go/pkg/mod/github.com
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go build -o /app/dcr_api && chmod +x /app/dcr_api


FROM debian:12.4
RUN useradd dcrapi
RUN apt-get update && apt-get install -y ca-certificates
COPY ./conf /usr/local/dcr_conf

WORKDIR /app

RUN mkdir /app/data
RUN chown -R dcrapi:dcrapi /app/data
COPY --from=builder --chown=dcrapi:dcrapi /app/dcr_api /usr/local/bin/dcr_api
USER dcrapi
ENTRYPOINT dcr_api
