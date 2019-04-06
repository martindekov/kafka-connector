#!/bin/sh
set -e

if [ -z "$NAMESPACE" ]; then
    NAMESPACE="functions"
fi

docker push $NAMESPACE/kafka-connector:$TAG

