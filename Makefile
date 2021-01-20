.PHONY: build-all build-master build-module images push-all push-master push-module

build-all: build-master build-module

build-master:
	docker-compose \
		--file master/docker-compose.yml \
		build \
		master

build-module:
	docker-compose \
		--file modules/base/docker-compose.yml \
		build \
		base

images push-all: push-master push-module

push-master: build-master
	docker-compose \
		--file master/docker-compose.yml \
		push \
		master

push-module: build-module
	docker-compose \
		--file modules/base/docker-compose.yml \
		push \
		base
