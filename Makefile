default: generate

.PHONY: generate
generate:
	mkdir -p dist
	cp -r png dist/
	go run ./cmd/generate/

