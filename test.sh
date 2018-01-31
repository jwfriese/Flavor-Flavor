#!/bin/bash

set -e

echo 'Resetting the database...'
PGPASSWORD=moreflava psql -Uflavor-flavor -hlocalhost -dpostgres -f ./scripts/reset.sql

echo 'Running tests...'
ginkgo
