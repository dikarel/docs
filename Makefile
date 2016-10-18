build: bin/docs

clean:
	@rm -rf bin/*

test:
	@go test

bin/docs: *.go
	@go build -o bin/docs
