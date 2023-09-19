build:
	@go build -o bin/fileserver
run:	build
	@./bin/fileserver
test:
	@go test ./... -v
