# Builder
FROM golang:1.20.6-alpine3.18 AS builder

ENV GO111MODULE=on

RUN apk update && apk --no-cache add git make

ADD . /tanshogyo

WORKDIR /tanshogyo/services/user

RUN make setup
RUN make build

# Runner
FROM alpine:3.18

RUN apk add --no-cache bash curl

WORKDIR /tanshogyo-user

COPY --from=builder /tanshogyo/services/user/bin/app /tanshogyo-user/bin/app
COPY --from=builder /tanshogyo/services/user/configs /tanshogyo-user/configs
COPY --from=builder /tanshogyo/services/user/scripts /tanshogyo-user/scripts

HEALTHCHECK --interval=1m --timeout=10s --start-period=30s --retries=3 CMD ["curl", "-f", "http://localhost"]

ENTRYPOINT ["bash", "./scripts/docker-entrypoint.sh", "./bin/app"]
