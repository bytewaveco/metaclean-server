#!/bin/bash

set -e

source meta.env

docker run --rm -it -p "3333:3333" -v "$(pwd)/src:/home/server" --name ${META_IMAGE_DEV} ${META_IMAGE_DEV}

set +e
