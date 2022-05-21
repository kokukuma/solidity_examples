package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/kokukuma/web3/solidity/auction"
	"github.com/kokukuma/web3/solidity/coin"
	"github.com/kokukuma/web3/solidity/sample"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"
	accountAddress = "7E532D0E80480eCF1b566920752840bc53556366"
)

type Deployer func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error)

func main() {
	if err := Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := sample.DeploySample(auth, backend)
		return contractAddress, tx, err
	}); err != nil {
		log.Fatalf("failed to deploy sample: %v", err)
	}

	if err := Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := coin.DeployCoin(auth, backend)
		return contractAddress, tx, err
	}); err != nil {
		log.Fatalf("failed to deploy coin: %v", err)
	}

	if err := Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := auction.DeployAuction(auth, backend, big.NewInt(2592000), common.HexToAddress(accountAddress))
		return contractAddress, tx, err
	}); err != nil {
		log.Fatalf("failed to deploy auction: %v", err)
	}
}

func Deploy(deployFunc Deployer) error {
	ctx := context.Background()

	client, err := ethclient.Dial(rpcProviderURL)
	if err != nil {
		return err
	}
	defer client.Close()

	address := common.HexToAddress(fmt.Sprintf("0x%s", accountAddress))
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return err

	}
	fmt.Printf("Balance: %d\n", balance)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Gas Price: %d\n", gasPrice)

	auth, err := GetAuth()
	if err != nil {
		return err
	}
	auth.GasPrice = gasPrice

	contractAddress, tx, err := deployFunc(auth, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	fmt.Printf("Contract Address: %s", contractAddress.Hex())

	return nil
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
