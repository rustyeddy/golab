subdirs = webster
mainfiles = webster/main.go webster/commands.go

src = /Users/rusty/src/rustyeddy.com
dst = /srv/www/rustyeddy.com

default: build test

build:
	go build .

test:
	go test .

install:
	go install

debug:
	/usr/local/bin/dlv debug --headless --listen=:2345 --api-version=2 webster/main.go gen --src ./etc/site.simple --dst /srv/www/_site.simple
