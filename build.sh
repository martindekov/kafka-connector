#!/bin/sh

if [ -z "$TRAVIS_TAG" ]; then
    TRAVIS_TAG="latest"
fi

docker build -t functions/kafka-connector:$TRAVIS_TAG .
(cd setup && docker service rm kafka_connector ; docker stack deploy kafka -c connector-swarm.yml)
