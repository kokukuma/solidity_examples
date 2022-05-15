package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kokukuma/web3/solidity/sample"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"
)

func main() {
	client, err := ethclient.Dial(rpcProviderURL)
	if err != nil {
		fmt.Println("failed to get client", err)
		return
	}

	// abigen --abi contract.abi --pkg main --type contract --out contract.go
	instance, err := sample.NewSample(common.HexToAddress(contractAddress), client)
	if err != nil {
		fmt.Println("failed to get contract", err)
		return
	}
	fmt.Println(instance)

	// balance, err := instance.BalanceOf(nil, common.HexToAddress(accountAddress))
	// if err != nil {
	// 	fmt.Println("failed to get balanceOf", err)
	// 	return
	// }
	// fmt.Println(balance)
	//
	// owner, err := instance.OwnerOf(nil, big.NewInt(1))
	// if err != nil {
	// 	fmt.Println("failed to get balanceOf", err)
	// 	return
	// }
	// fmt.Println(owner)
}
