#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

export GO111MODULE=on
export GOFLAGS="-v"
export SKIP_KNOWN_FAIL="1"

# Launch the arguments with the configured environment.
exec "$@"
