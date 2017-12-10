default: build 

lint:
	go fmt ./src/
	go vet ./src/
	gometalinter --deadline=15s ./src/...

dep:
	go get -v ./src/

build: test 
	CGO_ENABLED=0 go build -o er-task ./src/
test: dep
	docker-compose -f test.yml  up --build --force-recreate --remove-orphans  -d  
	sleep 10
	docker exec -i `docker ps --filter='ancestor=postgres' -q` psql -U postgres < ./sql/createdb.sql
	docker exec -i `docker ps --filter='ancestor=postgres' -q` psql -U postgres carrental < ./sql/db.dump
	go test -v ./...
	docker-compose -f test.yml down || true

	