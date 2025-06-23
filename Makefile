build:
	go build ./cmd/server

run-console:
	go build ./cmd/console
	./console
	open chart.html

test:
	go test -v ./...

clean:
	go clean
	rm -f server
	rm -rf chart.html
	rm -rf console