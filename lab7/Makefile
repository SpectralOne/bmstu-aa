.PHONY: run bench

run: 
	go run src/main.go

install:
	go get -v github.com/brianvoe/gofakeit

bench:
	cd src/dict/ && go test -bench . -benchmem