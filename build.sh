#!/bin/bash

set -e

source meta.env

docker build -t ${META_IMAGE} .

set +e
