build:
	go build ./cmd/server

run-console:
	go build ./cmd/console
	./console

test:
	go test -v ./...

clean:
	go clean
	rm -f server