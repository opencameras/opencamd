.PHONY: vendor

SHELL=/bin/bash # Use bash syntax

gen/opencamera:
	swagger generate server -f ./spec/opencameras.yaml -t ./gen/opencamera -P models.User --exclude-main
	swagger generate client -f ./spec/opencameras.yaml -t ./gen/opencamera --skip-models
