# Builder
FROM golang:1.20.6-alpine3.18 AS builder

ENV GO111MODULE=on

RUN apk update && apk --no-cache add git make

ADD . /tanshogyo

WORKDIR /tanshogyo/services/transaction

RUN make setup
RUN make build

# Runner
FROM alpine:3.18

RUN apk add --no-cache bash curl

WORKDIR /tanshogyo-transaction

COPY --from=builder /tanshogyo/services/transaction/bin/app /tanshogyo-transaction/bin/app
COPY --from=builder /tanshogyo/services/transaction/configs /tanshogyo-transaction/configs
COPY --from=builder /tanshogyo/services/transaction/scripts /tanshogyo-transaction/scripts

HEALTHCHECK --interval=1m --timeout=10s --start-period=30s --retries=3 CMD ["curl", "-f", "http://localhost"]

ENTRYPOINT ["bash", "./scripts/docker-entrypoint.sh", "./bin/app"]
