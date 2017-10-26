all: pong

clean:
	rm -rf dist

deps:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/chuckpreslar/emission
	go get github.com/spf13/cobra/cobra
	go get github.com/mitchellh/colorstring
	go get github.com/hashicorp/go-multierror
	go get github.com/davecgh/go-spew/spew
	go get github.com/hashicorp/consul/api

pong: clean deps
	go build -o dist/usr/bin/pong github.com/solvent-io/pong/cli/pong

illumos: clean deps
	GOOS=solaris go build -o dist/illumos/usr/bin/pong github.com/solvent-io/pong/cli/pong
	
fmt:
	goimports -w .

.PHONY: deps fmt
