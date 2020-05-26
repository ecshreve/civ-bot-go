build:
	go build -o bin/civ-bot github.com/ecshreve/civ-bot-go/cmd/civ-bot-go

test:
	go test -v github.com/ecshreve/civ-bot-go/...