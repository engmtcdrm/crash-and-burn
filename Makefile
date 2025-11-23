.PHONY: build runexe run

build:
	echo "Size before build:"; ls -la |grep crash-and-burn; ls -lh |grep crash-and-burn; echo "\n\nSize after build:"; CGO_ENABLED=0 go build --ldflags "-s -w" -o crash-and-burn; strip crash-and-burn; ls -la |grep crash-and-burn; ls -lh |grep crash-and-burn

runexe:
	./crash-and-burn $(ARGS)

run:
	go run . $(ARGS)

test:
	go test -timeout 30s ./...

testv:
	go test -timeout 30s -v ./...
