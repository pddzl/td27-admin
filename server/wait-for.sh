#!/bin/sh

set -e

until nc -w2 mysql 3306 && nc -w2 redis 6379
do
  echo "mysql or redis is unavailable - sleeping" >&2
  sleep 1
done

exec "$@"
