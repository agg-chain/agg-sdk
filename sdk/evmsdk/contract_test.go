package evmsdk

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"
)

func TestCreateContract(t *testing.T) {
	// 0xD64229dF1EB0354583F46e46580849B1572BB56d
	privateKey, err := crypto.HexToECDSA("5f9468d68c0fa82f78d8fe774c6d860ca8ff4a2f4f485603144183666dd8bf1d")
	if err != nil {
		log.Fatal(err)
	}

	// See ./test_sol/storage.sol for helper functions used in this file.
	data, err := hex.DecodeString("608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea2646970667358221220322c78243e61b783558509c9cc22cb8493dde6925aa5e89a08cdf6e22f279ef164736f6c63430008120033")
	if err != nil {
		panic(err)
	}
	toAddress := common.HexToAddress("0x")
	contractAddress, err := CreateContract(privateKey, &toAddress, data)
	if err != nil {
		return
	}
	println("Contract address: " + contractAddress)
}

func TestCallNoGasContract(t *testing.T) {
	// retrieve code: 0x2e64cec1
	inputBytes, err := hexutil.Decode("0x2e64cec1")
	if err != nil {
		panic(err)
	}
	// contract address: 0xBd770416a3345F91E4B34576cb804a576fa48EB1
	contractAddress := common.HexToAddress("0xbd770416a3345f91e4b34576cb804a576fa48eb1")
	contract, err := CallContract(&contractAddress, inputBytes)
	if err != nil {
		panic(err)
	}
	result := new(big.Int)
	result.SetString(hexutil.Encode(contract)[2:], 16)
	println(result.String())
}

func TestExecuteNeedGasContract(t *testing.T) {
	// 0xD64229dF1EB0354583F46e46580849B1572BB56d
	privateKey, err := crypto.HexToECDSA("5f9468d68c0fa82f78d8fe774c6d860ca8ff4a2f4f485603144183666dd8bf1d")
	if err != nil {
		log.Fatal(err)
	}

	data, err := hex.DecodeString("6057361d0000000000000000000000000000000000000000000000000000000000000001")
	if err != nil {
		panic(err)
	}
	// contract address: 0xBd770416a3345F91E4B34576cb804a576fa48EB1
	contractAddress := common.HexToAddress("0xbd770416a3345f91e4b34576cb804a576fa48eb1")
	contract, err := ExecuteContract(privateKey, &contractAddress, data)
	if err != nil {
		panic(err)
	}
	println(contract.Result.DeliverTx.Data)
}

func TestExecuteNoGasContract(t *testing.T) {
	// 0xD64229dF1EB0354583F46e46580849B1572BB56d
	privateKey, err := crypto.HexToECDSA("5f9468d68c0fa82f78d8fe774c6d860ca8ff4a2f4f485603144183666dd8bf1d")
	if err != nil {
		log.Fatal(err)
	}

	// retrieve code: 0x2e64cec1
	data, err := hex.DecodeString("2e64cec1")
	if err != nil {
		panic(err)
	}
	// contract address: 0xBd770416a3345F91E4B34576cb804a576fa48EB1
	contractAddress := common.HexToAddress("0xbd770416a3345f91e4b34576cb804a576fa48eb1")
	contract, err := ExecuteContract(privateKey, &contractAddress, data)
	if err != nil {
		panic(err)
	}
	println(contract.Result.DeliverTx.Data)
}
