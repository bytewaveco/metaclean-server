#!/bin/bash

set -e

source meta.env

docker build -t ${META_IMAGE}-dev -f Dockerfile.dev .

set +e
