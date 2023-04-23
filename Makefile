.PHONEY: bin
build:
	go build -o bin/movie-crud-service

run: build
	./bin/movie-crud-service
