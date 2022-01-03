#!/bin/bash

set -e

WORKING_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
WORKING_DIR=${WORKING_DIR}/src

if [[ ! -z ${META_S_PATH} ]]; then
    WORKING_DIR=${META_S_PATH}
fi

echo "Setting up project for development..."

( cd ${WORKING_DIR} && go mod download )
( cd ${WORKING_DIR} && go mod vendor )

echo "✨ You're good to Go!"

set +e
