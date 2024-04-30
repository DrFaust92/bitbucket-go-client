default: patch generate fmt build

fmt:
	go fmt

build:
	go build

patch:
	go install github.com/evanphx/json-patch/cmd/json-patch@v5
	wget -q --output-document - https://api.bitbucket.org/swagger.json | "$$(go env GOPATH)/bin/json-patch" -p api/swagger.json-patch > api/swagger.output.json

generate: patch
	swagger-codegen generate -i api/swagger.output.json -l go -c swagger.conf --additional-properties packageName=bitbucket --template-dir swagger-template
