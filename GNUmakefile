default: patch generate postprocess fmt build

fmt:
	go fmt

build:
	go build

patch:
	go install github.com/evanphx/json-patch/cmd/json-patch@v5
	wget -q --output-document - https://api.bitbucket.org/swagger.json | "$$(go env GOPATH)/bin/json-patch" -p api/swagger.json-patch > api/swagger.output.json

generate: patch
	swagger-codegen generate -i api/swagger.output.json -l go -c swagger.conf --additional-properties packageName=bitbucket

# swagger-codegen emits tags that break write requests (see #41):
# - `type` is required in the spec, so it has no omitempty and an unset
#   value is sent as "type":"", which Bitbucket rejects on POST/PUT;
# - time.Time fields use omitempty, which never omits a struct, so unset
#   timestamps are sent as "0001-01-01T00:00:00Z". omitzero (Go 1.24+) omits
#   them via time.Time.IsZero.
postprocess:
	perl -pi -e 's/`json:"type"`/`json:"type,omitempty"`/; s/(time\.Time\s+`json:"[^",]+),omitempty"/$$1,omitzero"/' model_*.go

.PHONY: default fmt build patch generate postprocess
