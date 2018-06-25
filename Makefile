.PHONY: licensezero test

LDFLAGS=-X main.Rev=$(git rev-parse --short HEAD)

licensezero:
	go build -o licensezero -ldflags "$(LDFLAGS)"

test: licensezero
	go test

build:
	gox -os="linux darwin windows freebsd" -arch="i386 amd64" -output="licensezero-{{.OS}}-{{.Arch}}" -ldflags "$(LDFLAGS)" -verbose
