#!/usr/bin/env bash

set -e

scriptdir=$(dirname "$(readlink -f "$BASH_SOURCE")")

if [ -n "$PRODUCT_DATASOURCES_GORM_CONNECTIONS_MAIN_HOST" ]; then
    $scriptdir/wait-for-it.sh -t 60 "$PRODUCT_DATASOURCES_GORM_CONNECTIONS_MAIN_HOST:$PRODUCT_DATASOURCES_GORM_CONNECTIONS_MAIN_PORT"
fi

exec "$@"
