.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./redirick

.PHONY: build
build:
	docker build -t ghcr.io/dyrector-io/redirick:2 .

.PHONY: push
push:
	docker push ghcr.io/dyrector-io/redirick:2

.PHONY: clean
clean:
	rm ./redir || true

all: clean compile build push
