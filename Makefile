run:
	go run gymshark-knapsack/cmd

test:
	go test -v ./...

lint:
	golangci-lint run