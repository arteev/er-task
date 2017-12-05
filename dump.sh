#!/bin/sh
for container_id in $(docker ps --filter='ancestor=postgres' -q)
  do
    echo "Clear sample data..."
    docker exec -i $container_id psql -U postgres carrental < ./_sql/clear.db.sql
    echo "Dump..."
    docker exec -i $container_id pg_dump carrental -U postgres > ./_sql/db.dump       
  done
echo "DONE"
