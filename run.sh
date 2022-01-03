#!/bin/bash

set -e

source meta.env

docker run --rm -it -p "3333:3333" --name ${META_IMAGE} ${META_IMAGE}

set +e
