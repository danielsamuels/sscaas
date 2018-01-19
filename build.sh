#!/usr/bin/env bash

BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

REPO_NAME="manafont"
IMAGE_NAME="sscaas"
IMAGE_VERSION="1.0.0"

##
## Invoke Docker Build
##

docker build \
    -t "${REPO_NAME}/${IMAGE_NAME}:${IMAGE_VERSION}" \
    -f "${BASE}/Dockerfile" \
    "${BASE}"

##
## Invoke Docker Push
##

# docker push "${REPO_NAME}/${IMAGE_NAME}:${IMAGE_VERSION}"
