#!/bin/bash

set -e

until nc -w2 -t mysql 3306 && nc -w2 -t redis 6379
do
  echo "mysql or redis is unavailable - sleeping" >&2
  sleep 1
done

exec "$@"
