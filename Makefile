build:
	go build -o bin/civ-bot github.com/ecshreve/civ-bot-go/cmd/civ-bot-go

run-only:
	bin/civ-bot

run: build run-only

test:
	go test -v github.com/ecshreve/civ-bot-go/...