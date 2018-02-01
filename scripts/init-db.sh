#!/bin/bash
set -e

spinUpContainer() {
  if [ ! $(docker ps -q -f name=flavor-flavor-postgres) ]; then
    if [ "$(docker ps -aq -f status=exited -f name=flavor-flavor-postgres)" ]; then
      docker rm flavor-flavor-postgres
    fi

    docker run --name flavor-flavor-postgres -e POSTGRES_USER=flavor-flavor -e POSTGRES_PASSWORD=moreflava -d -p 5432:5432 postgres:9.5
    until pg_isready -hlocalhost -p5432; do
    >&2 echo "Waiting for Postgres to come up..."
    sleep 1
    done
  fi
}

initializeDatabase() {
  PGPASSWORD=moreflava psql -Uflavor-flavor -hlocalhost -dpostgres -f ./scripts/create.sql
  PGPASSWORD=moreflava psql -Uflavor-flavor -hlocalhost -dpostgres -f ./scripts/reset.sql
}

spinUpContainer
initializeDatabase
