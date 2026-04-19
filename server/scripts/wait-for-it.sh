#!/bin/sh

set -e

until nc -w2 pgsql 5432
do
  echo "pgsql is unavailable - sleeping" >&2
  sleep 1
done

exec "$@"
