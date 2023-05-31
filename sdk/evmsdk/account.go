package evmsdk

import (
	"agg-sdk/model"
	"agg-sdk/sdkconfig"
	"agg-sdk/util"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func GetNonce(address string) (uint64, error) {
	decoded, err := abciQuery("get_nonce=" + address)
	if err != nil {
		return 0, err
	}

	nonce := new(big.Int)
	nonce.SetString(string(decoded), 16)
	return nonce.Uint64(), nil
}

func GetBalance(address string) (*big.Int, error) {
	decoded, err := abciQuery("get_balance=" + address)
	if err != nil {
		return nil, err
	}

	balance := new(big.Int)
	balance.SetString(string(decoded[2:]), 16)
	return balance, nil
}

func GetTx(address, start, end string) (string, error) {
	decoded, err := abciQuery("get_tx=" + address + "_" + start + "_" + end)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func Transfer(privateKey *ecdsa.PrivateKey, toAddress *common.Address, value *big.Int) (string, error) {

	nonce, err := GetNonce(GetAddressHex(privateKey))
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000) // in units
	gasPrice := big.NewInt(20000)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     []byte{},
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(sdkconfig.ChainId), privateKey)
	if err != nil {
		return "", err
	}

	// get hex data
	txData, err := signedTx.MarshalBinary()
	if err != nil {
		return "", err
	}
	txDataHex := hexutil.Encode(txData)
	fmt.Printf("tx data: %s\n", txDataHex)

	txInAgg := &types.Transaction{}
	rawTxBytes, err := hex.DecodeString(txDataHex[2:])
	if err != nil {
		return "", err
	}

	err = txInAgg.UnmarshalBinary(rawTxBytes)
	if err != nil {
		return "", err
	}

	res := util.Get(fmt.Sprintf("%s/broadcast_tx_commit?tx=\""+txDataHex+"\"", sdkconfig.RpcUrl))
	broadcastResponse := model.BroadcastTxCommitResponse{}
	err = json.Unmarshal([]byte(res), &broadcastResponse)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(broadcastResponse.Result.DeliverTx.Data.(string))
	if err != nil {
		return "", err
	}
	return hexutil.Encode(decoded), nil
}
