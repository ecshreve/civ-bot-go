DIRECTORIES=$(sort $(dir $(wildcard pkg/*/*/)))
MOCKS=$(foreach x, $(DIRECTORIES), mocks/$(x))

out:
	@echo $(DIRECTORIES)

build:
	go build -o bin/civ-bot github.com/ecshreve/civ-bot-go/cmd/civ-bot-go

run-only:
	bin/civ-bot

run: build run-only

test: | mocks
	go test github.com/ecshreve/civ-bot-go/...

testc: | mocks
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/civ-bot-go/...

testv: | mocks
	go test -v github.com/ecshreve/civ-bot-go/...

clean-mocks:
	rm -rf mocks

mocks: $(MOCKS)
  
$(MOCKS): mocks/% : %
	mockery -output=$@ -dir=$^ -all