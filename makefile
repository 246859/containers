
lint:
	go vet ./...

test:
	make lint
	go test -v ./...

build:
	make test
	go build -v ./...


.PHONY: lint \
		test \
		build