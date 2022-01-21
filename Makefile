build: clean
	go build github.com/ajean-zd/gameoflife/cmd/gameoflife

clean:
	rm -f gameoflife

test:
	go test ./internal/...
