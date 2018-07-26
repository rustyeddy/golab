verbose = 
cleandirs = tmp

build:
	go build 

install:
	go install $(verbose) 

test:
	go test $(verbose)
	make -C tests test

coverage:
	go test -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o docs/coverage.html 

clean:
	go clean
	rm -rf $(cleandirs)
	rm -rf *~

docs/godocs.txt: $(src)

.PHONY: tests run install build coverage clean
