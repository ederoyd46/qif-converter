.PHONY: build

buid:
	go build .

release:
	go build -ldflags "-w" .
