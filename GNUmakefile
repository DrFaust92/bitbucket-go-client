default: generate fmt build

fmt:
	go fmt

build:
	go build

generate:
	swagger-codegen generate -i https://api.bitbucket.org/swagger.json -l go -c swagger.conf --additional-properties packageName=bitbucket
