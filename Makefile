default: build 

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s ./...

dep:
	go get -v 

build: test 
	CGO_ENABLED=0 go build ${LDFLAGS}
test: dep
	docker-compose -f test.yml  up --build --force-recreate --remove-orphans  -d  
	sleep 5
	docker exec -i `docker ps --filter='ancestor=postgres' -q` psql -U postgres < ./_sql/createdb.sql
	docker exec -i `docker ps --filter='ancestor=postgres' -q` psql -U postgres carrental < ./_sql/db.dump
	go test -v ./...
	docker-compose -f test.yml down || true

	