
.PHONY: install-dep test dev

install-dep:
	go get github.com/pm5/wat

test:
	cd test; DEBUG=* go test

dev:
	wat 'make test' -- Makefile *.go
