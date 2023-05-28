package evmsdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"
)

func TestGetBalance(t *testing.T) {
	res, err := GetBalance("0xD64229dF1EB0354583F46e46580849B1572BB56d")
	//res, err := GetBalance("0x27a01491d86F3F3b3085a0Ebe3F640387DBdb0EC")
	if err != nil {
		panic(err)
	}
	println(res.String())
}

func TestTransfer(t *testing.T) {
	// 0xD64229dF1EB0354583F46e46580849B1572BB56d
	privateKey, err := crypto.HexToECDSA("5f9468d68c0fa82f78d8fe774c6d860ca8ff4a2f4f485603144183666dd8bf1d")
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x27a01491d86F3F3b3085a0Ebe3F640387DBdb0EC")

	value := big.NewInt(0.001 * 1e18) // in wei (1 eth)

	txHash, err := Transfer(privateKey, &toAddress, value)
	if err != nil {
		return
	}
	fmt.Printf("txHash hash: %s\n", txHash)

}

func TestGetNonce(t *testing.T) {
	println(GetNonce("0xD64229dF1EB0354583F46e46580849B1572BB56d"))
}
