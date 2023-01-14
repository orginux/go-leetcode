fmt:
	go fmt ./...

test-all:
	find . -maxdepth 1 -type d -not -path './.*' -not -path . -exec go test -cover {}  \;

