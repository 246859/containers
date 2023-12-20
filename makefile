
lint:
	go vet ./...

test:
	make lint
	go test -v -count=1 ./...

build:
	make test
	go build -v ./...


.PHONY: lint \
		test \
		build