package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/kokukuma/web3/solidity/sample"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"

	contractAddress = "8b6F44eB598e240d8608b42cc04EBdBfA360BB09"
	accountAddress  = "ee6983ACaCe98e766E5b78f0533EA74c6e734dE5"
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

	//
	gasPrice, err := GasPrice()
	if err != nil {
		fmt.Println(err)
		return
	}

	// auth
	auth, err := GetAuth()
	if err != nil {
		fmt.Println(err)
		return
	}
	auth.GasPrice = gasPrice

	// opts *bind.TransactOpts, x *big.Int
	tx, err := instance.Set(auth, big.NewInt(1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Set deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	val, err := instance.Get(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func GetAuth() (*bind.TransactOpts, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	keystore, err := os.Open(os.Getenv("KEYSTORE"))
	if err != nil {
		return nil, err
	}
	defer keystore.Close()

	keystorepass := os.Getenv("KEYSTOREPASS")
	return bind.NewTransactor(keystore, keystorepass)
}

func GasPrice() (*big.Int, error) {
	ctx := context.Background()

	client, err := ethclient.Dial(rpcProviderURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	address := common.HexToAddress(accountAddress)
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err

	}
	fmt.Printf("Balance: %d\n", balance)

	return client.SuggestGasPrice(context.Background())
}
