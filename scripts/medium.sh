#!/bin/bash

set -e
set -u
set -o pipefail

docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 elasticsearch
DOCKER_HOST=${DOCKER_HOST-}
if [[ -z "${DOCKER_HOST}" ]]; then
  SNAP_ELASTICSEARCH_HOST="127.0.0.1"
else
  SNAP_ELASTICSEARCH_HOST=$(docker inspect -f '{{ .NetworkSettings.IPAddress }}' elasticsearch)
fi

export SNAP_ELASTICSEARCH_HOST
_info "Elasticsearch Host: ${SNAP_ELASTICSEARCH_HOST}"

_info "Waiting for elasticsearch docker container"
while ! curl -sG "http://${SNAP_ELASTICSEARCH_HOST}:9200/_status" > /dev/null 2>&1; do
  sleep 1
  echo -n "."
done
echo

UNIT_TEST="go_test"
test_unit
