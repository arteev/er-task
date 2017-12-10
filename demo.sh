#!/bin/sh
# запуск приложения с демо базой

docker-compose -f demo.yml  up --build --force-recreate --remove-orphans -d 

sleep 10
echo "Creating sample database..."
for container_id in $(docker ps --filter='ancestor=postgres' -q)
  do
    docker exec -i $container_id psql -U postgres < ./sql/createdb.sql
    docker exec -i $container_id psql -U postgres carrental < ./sql/db.dump
    docker exec -i $container_id psql -U postgres carrental < ./sql/example.db.sql
  done
echo "Creating sample database...DONE"

#scale 
#docker-compose -f demo.yml scale web=3