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
	"github.com/kokukuma/web3/solidity/auction"
	"github.com/kokukuma/web3/solidity/coin"
	"github.com/kokukuma/web3/solidity/sample"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"

	contractAddress        = "8b6F44eB598e240d8608b42cc04EBdBfA360BB09"
	coinContractAddress    = "9a09934744214B88929c15DA37e08c47d7b8DC2a"
	auctionContractAddress = "4A19B6f557A7e314E19d72483dadb4Fac468101A"
	accountAddress         = "7E532D0E80480eCF1b566920752840bc53556366"
	metamaskAddress        = "ee6983ACaCe98e766E5b78f0533EA74c6e734dE5"
)

func main() {
	ctx := context.Background()

	client, err := ethclient.Dial(rpcProviderURL)
	if err != nil {
		fmt.Println("failed to get client", err)
		return
	}

	auth, err := GetAuth()
	if err != nil {
		fmt.Println("failed to get auth", err)
		return
	}
	defer client.Close()

	// if err := Sample(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := CoinMint(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := CoinSend(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := CoinGet(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }

	if err := AuctionGetHighestBid(ctx, client, auth); err != nil {
		fmt.Println(err)
	}

	// if err := AuctionBid(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }

	// if err := AuctionWithdraw(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }

	if err := AuctionEnd(ctx, client, auth); err != nil {
		fmt.Println(err)
	}
}

func AuctionEnd(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := auction.NewAuction(common.HexToAddress(auctionContractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to new auction: %s", err.Error())
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %s", err.Error())
	}
	auth.GasPrice = gasPrice
	fmt.Printf("value: %s, gas: %s\n", auth.Value.String(), auth.GasPrice.String())

	tx, err := instance.AuctionEnd(auth)
	if err != nil {
		// if `require` is not satisfied, this error will happen.
		return fmt.Errorf("failed to end: %s", err.Error())
	}
	fmt.Printf("End! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	return nil
}

func AuctionWithdraw(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := auction.NewAuction(common.HexToAddress(auctionContractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to new auction: %s", err.Error())
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %s", err.Error())
	}
	auth.GasPrice = gasPrice
	fmt.Printf("value: %s, gas: %s\n", auth.Value.String(), auth.GasPrice.String())

	tx, err := instance.Withdraw(auth)
	if err != nil {
		// if `require` is not satisfied, this error will happen.
		return fmt.Errorf("failed to withdraw: %s", err.Error())
	}
	fmt.Printf("Withdraw! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	return nil
}

func AuctionGetHighestBid(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := auction.NewAuction(common.HexToAddress(auctionContractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to new auction: %s", err.Error())
	}
	val, err := instance.HighestBid(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	})
	if err != nil {
		return fmt.Errorf("failed to get highest bid: %s", err.Error())
	}
	fmt.Printf("Current HighestBid: %s\n", val.String())
	return nil
}

func AuctionBid(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := auction.NewAuction(common.HexToAddress(auctionContractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to new auction: %s", err.Error())
	}
	// set value and gas
	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %s", err.Error())
	}
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(3)
	fmt.Printf("value: %s, gas: %s\n", auth.Value.String(), auth.GasPrice.String())

	tx, err := instance.Bid(auth)
	if err != nil {
		// if `require` is not satisfied, this error will happen.
		return fmt.Errorf("failed to bid: %s", err.Error())
	}
	fmt.Printf("Minted! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	return nil
}

func CoinMint(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := coin.NewCoin(common.HexToAddress(coinContractAddress), client)
	if err != nil {
		return err
	}
	tx, err := instance.Mint(auth, common.HexToAddress(accountAddress), big.NewInt(100))
	if err != nil {
		return err
	}
	fmt.Printf("Minted! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	return nil
}

func CoinSend(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := coin.NewCoin(common.HexToAddress(coinContractAddress), client)
	if err != nil {
		return err
	}
	tx, err := instance.Send(auth, common.HexToAddress(metamaskAddress), big.NewInt(10))
	if err != nil {
		return err
	}
	fmt.Printf("Sended! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	return nil
}

func CoinGet(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := coin.NewCoin(common.HexToAddress(coinContractAddress), client)
	if err != nil {
		return err
	}
	balance, err := instance.Get(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	}, common.HexToAddress(metamaskAddress))
	if err != nil {
		return err
	}
	fmt.Printf("Balance: %v", balance)
	return nil
}

func Sample(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := sample.NewSample(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return err
	}
	auth.GasPrice = gasPrice

	// opts *bind.TransactOpts, x *big.Int
	tx, err := instance.Set(auth, big.NewInt(1))
	if err != nil {
		return err
	}
	fmt.Printf("Set deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	val, err := instance.Get(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	})
	if err != nil {
		return err
	}
	fmt.Println(val)

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

func GasPrice(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err

	}
	fmt.Printf("Balance: %d\n", balance)

	return client.SuggestGasPrice(context.Background())
}
