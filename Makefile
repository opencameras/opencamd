.PHONY: vendor

SHELL=/bin/bash # Use bash syntax

bootstrap-opencamd:
	rm -rf ./gen/opencamera/*
	swagger generate server -f ./spec/opencameras.yaml -t ./gen/opencamera -P models.Principal
	swagger generate client -f ./spec/opencameras.yaml -t ./gen/opencamera --skip-models

regen/opencamera:
	rm -rf ./gen/opencamera/*
	swagger generate server -f ./spec/opencameras.yaml -t ./gen/opencamera -P models.Principal --exclude-main
	swagger generate client -f ./spec/opencameras.yaml -t ./gen/opencamera --skip-models
