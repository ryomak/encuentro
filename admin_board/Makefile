
.PHONY: all
all: run

.PHONY: run
run:
	npm run serve

.PHONY: run-local
run-local:
	USE_LOCAL_SERVER=1 npm run serve

.PHONY: build
build:
	npm run build
	# find ./dist -name *.map | xargs rm

.PHONY: deps
deps:
	npm install
