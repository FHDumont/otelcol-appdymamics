SHELL = /bin/bash

.PHONY: build
build:
	builder --config=builder-config.yaml 

.PHONY: run
run: build
	set -o allexport
	./otelcol-dev/otelcol-dev --config=config.yaml

