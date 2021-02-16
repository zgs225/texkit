pack:
	pkger -include github.com/zgs225/texkit:/templates -o generator

build: pack
	go build .
