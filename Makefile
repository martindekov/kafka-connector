TAG?=latest
NAMESPACE?=functions
.PHONY: build

build:
	docker build -t $(NAMESPACE)/kafka-connector:$(TAG) .
push:
	docker push $(NAMESPACE)/kafka-connector:$(TAG)
