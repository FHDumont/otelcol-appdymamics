#!/bin/bash

CONTAINER_VERSION=1.0
DOCKER_DOMAIN=fernandodumont

PUSH_MA=false
PUSH_OTEL=false

for ARGUMENT in "$@"
do
    if [ "$ARGUMENT" = "otel" ]; then
        BUILD_OTEL=true
    elif [ "$ARGUMENT" = "ma" ]; then
        PUSH_MA=true
    elif [ "$ARGUMENT" = "all" ]; then
        BUILD_OTEL=true
        PUSH_MA=true
    fi
done

export DOCKER_BUILDKIT=1;

if [ $BUILD_OTEL == true ];then
    echo ""
    echo "==> OTEL"
    ( 
        docker buildx build --pull --push -t ${DOCKER_DOMAIN}/otelcol-appdynamics:${CONTAINER_VERSION} -o type=image --platform=linux/arm64,linux/amd64 -f otelcol/Dockerfile ..
    )
fi

if [ $PUSH_MA == true ];then
    echo ""
    echo "==> Machine Agent"
    ( 
        docker buildx build --pull --push -t ${DOCKER_DOMAIN}/machine-agent:${CONTAINER_VERSION} -o type=image --platform=linux/arm64,linux/amd64 -f machine-agent/Dockerfile .
    )
fi

