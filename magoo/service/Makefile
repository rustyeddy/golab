cmd = magoo

build:
	go build -o $(cmd) -v .

test:
	go test -v

vet:
	go vet -v .

install: $(cmd)
	go install -o $(cmd)

run:
	go run -v main.go
