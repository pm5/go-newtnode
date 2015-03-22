
.PHONY: test dev

test:
	cd test; DEBUG=* go test

dev:
	watch 'cd test; go test'
