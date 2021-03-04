all: init build
.PHONY: all

build: pack
	go build .

init:
	if ! which pkger > /dev/null; then go get github.com/markbates/pkger/cmd/pkger; fi

pack:
	pkger -include github.com/zgs225/texkit:/templates -o generator

.PHONY: clean
clean:
	-rm texkit

