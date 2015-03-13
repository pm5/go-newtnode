
.PHONY: test watch

test:
	cd test; DEBUG=* go test

watch:
	watch 'cd test; go test'
