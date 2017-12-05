default: build 

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s ./...

build: #test 
	CGO_ENABLED=0 go build ${LDFLAGS}
test: 
	go test -v ./...

	