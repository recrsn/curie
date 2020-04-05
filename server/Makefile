VERSION := $(shell git describe --tags --always --first-parent)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X github.com/curie/server/app.Version=$(VERSION)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

build: compile test

run: stop
	./curie

start: stop
		./curie& \
 		echo $$! >  curie.pid

stop:
	(kill $(file < curie.pid)  || true)
	rm -f curie.pid

compile:
	go build -o curie $(GOFLAGS) .

test:
	go test $(GOFLAGS) './...'