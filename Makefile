PORT = 5000

all: build

.PHONY: build
build:  ## make contract and abi
	solc --abi --bin sol/*.sol -o bin --overwrite

	abigen --abi="bin/SimpleStorage.abi" --bin="bin/SimpleStorage.bin" --pkg=sample --out="sample/sample.go"
	abigen --abi="bin/Coin.abi" --bin="bin/Coin.bin" --pkg=coin --out="coin/coin.go"

.PHONY: run
run:  ## run contract
	go run cmd/run/run.go

.PHONY: deploy
deploy: ## deploy contract
	go run cmd/deploy/deploy.go
