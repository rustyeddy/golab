cmd = upw

$(cmd): main.go ../*.go
	go build -o $(cmd)

run:
	go run *.go

runv:
	go run -v *.go

clean:
	go clean 
	rm *~
