#!/bin/sh

if [[ "${1#-}" != "$1" ]]; then
    set -- influxdb-admin "$@"
fi

exec "$@"
