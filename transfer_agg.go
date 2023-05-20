package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	AGG_CHAINID = 100
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/EmsfiRZ34lwj9uxyHC1lXktrBY7ivqBC")
	if err != nil {
		log.Fatal(err)
	}

	// 0xD64229dF1EB0354583F46e46580849B1572BB56d
	privateKey, err := crypto.HexToECDSA("5f9468d68c0fa82f78d8fe774c6d860ca8ff4a2f4f485603144183666dd8bf1d")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0.001 * 1e18) // in wei (1 eth)
	gasLimit := uint64(21000)         // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x27a01491d86F3F3b3085a0Ebe3F640387DBdb0EC")
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// get hex data
	txData, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx data: %s\n", hexutil.Encode(txData))

	// TODO send

	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Printf("tx hash: %s\n", signedTx.Hash().Hex())
}
