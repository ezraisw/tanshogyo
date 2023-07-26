# Builder
FROM golang:1.20.6-alpine3.18 AS builder

ENV GO111MODULE=on

RUN apk update && apk --no-cache add git make

ADD . /tanshogyo

WORKDIR /tanshogyo/services/product

RUN make setup
RUN make build

# Runner
FROM alpine:3.18

RUN apk add --no-cache bash curl

WORKDIR /tanshogyo-product

COPY --from=builder /tanshogyo/services/product/bin/app /tanshogyo-product/bin/app
COPY --from=builder /tanshogyo/services/product/configs /tanshogyo-product/configs
COPY --from=builder /tanshogyo/services/product/scripts /tanshogyo-product/scripts

HEALTHCHECK --interval=1m --timeout=10s --start-period=30s --retries=3 CMD ["curl", "-f", "http://localhost"]

ENTRYPOINT ["bash", "./scripts/docker-entrypoint.sh", "./bin/app"]
