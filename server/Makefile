VERSION := $(shell git describe --tags --always --first-parent)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X app.Version=$(VERSION)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

run: stop build
	./curie

start: stop build
		./curie& \
 		echo $$! >  curie.pid

stop:
	(kill $(file < curie.pid)  || true)
	rm -f curie.pid

compile:
	go build -o curie $(GOFLAGS) .

test:
	go test $(GOFLAGS)

build: compile test