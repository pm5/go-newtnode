
.PHONY: test watch

test:
	cd test; go test

watch:
	watch 'cd test; go test'
