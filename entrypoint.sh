#!/bin/bash

export DISCOVERY_ADDR=${DISCOVERY_ADDR}
export DISCOVERY_PORT=${DISCOVERY_PORT}

exec "$@"
