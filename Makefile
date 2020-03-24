run:
	go run .

test:
	go test -v -race .

.DEFAULT_GOAL=run
