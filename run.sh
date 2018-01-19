#!/usr/bin/env bash

# set -ex  # uncomment this for debug output

BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

REPO_NAME="manafont"
IMAGE_NAME="sscaas"
IMAGE_VERSION="1.0.0"

##
## Invoke Docker Run
##

docker run -d \
    --name "${IMAGE_NAME}" \
    -p "80:8080" \
    "${REPO_NAME}/${IMAGE_NAME}:${IMAGE_VERSION}"
