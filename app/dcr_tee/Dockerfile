FROM golang:1.22 as builder
WORKDIR /app
COPY ./attestation ./attestation
COPY ./encryption ./encryption

COPY ./github.com /go/pkg/mod/github.com
WORKDIR /app/attestation
RUN go build -o /app/gen_custom_token && chmod +x /app/gen_custom_token
WORKDIR /app/encryption
RUN go build -o /app/encrypt_tool && chmod +x /app/encrypt_tool

FROM quay.io/jupyter/scipy-notebook
RUN pip install google-cloud-storage
USER root
RUN apt-get update && apt-get -y install gnupg
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg && apt-get update -y && apt-get install google-cloud-sdk -y

COPY ./conf /usr/local/dcr_conf
COPY --from=builder /app/gen_custom_token /usr/local/bin/gen_custom_token
COPY --from=builder /app/encrypt_tool /usr/local/bin/encrypt_tool