.PHONY: run bench clean

run: main
	./main

main: src/main.go
	go build src/main.go

bench:
	cd src/levenshtein/ && go test -bench=. -benchmem

clean:
	rm -f main