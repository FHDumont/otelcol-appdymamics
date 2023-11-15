FROM alpine:3.16 as certs
RUN apk --update add ca-certificates

FROM golang:1.20.11-alpine3.18 as builder

WORKDIR /builder

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake 

COPY . .

RUN GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@v0.88.0
RUN make build

FROM alpine:latest as prep

RUN mkdir -p /tmp

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /builder/otelcol-dev/otelcol-dev /
COPY --from=builder /builder/config.yaml /etc/otel/config.yaml

ENTRYPOINT ["/otelcol-dev"]
CMD ["--config", "/etc/otel/config.yaml"]
EXPOSE 4317 55678 55679