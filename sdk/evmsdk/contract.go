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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func invokeContract(privateKey *ecdsa.PrivateKey, toAddress *common.Address, data []byte) (*model.BroadcastTxCommitResponse, error) {
	nonce, err := GetNonce(GetAddressHex(privateKey))
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(210000) // in units
	gasPrice := big.NewInt(20000)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       toAddress,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(sdkconfig.ChainId), privateKey)
	if err != nil {
		return nil, err
	}

	// get hex data
	txData, err := signedTx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	txDataHex := hexutil.Encode(txData)

	txInAgg := &types.Transaction{}
	rawTxBytes, err := hex.DecodeString(txDataHex[2:])
	if err != nil {
		return nil, err
	}

	err = txInAgg.UnmarshalBinary(rawTxBytes)
	if err != nil {
		return nil, err
	}

	res := util.Get(fmt.Sprintf("%s/broadcast_tx_commit?tx=\""+txDataHex+"\"", sdkconfig.RpcUrl))
	broadcastResponse := model.BroadcastTxCommitResponse{}
	err = json.Unmarshal([]byte(res), &broadcastResponse)
	if err != nil {
		return nil, err
	}

	return &broadcastResponse, nil
}

func CreateContract(privateKey *ecdsa.PrivateKey, toAddress *common.Address, data []byte) (string, error) {
	contractRes, err := invokeContract(privateKey, toAddress, data)
	if err != nil {
		return "", err
	}
	decoded, err := base64.StdEncoding.DecodeString(contractRes.Result.DeliverTx.Data.(string))
	if err != nil {
		return "", err
	}
	return "0x" + hex.EncodeToString(decoded), nil
}

func ExecuteContract(privateKey *ecdsa.PrivateKey, toAddress *common.Address, data []byte) (*model.BroadcastTxCommitResponse, error) {
	return invokeContract(privateKey, toAddress, data)
}

func CallContract(toAddress *common.Address, input []byte) ([]byte, error) {
	msg := ethereum.CallMsg{
		To:   toAddress,
		Data: input,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	decoded, err := abciQuery("eth_call=" + hexutil.Encode(msgBytes))
	if err != nil {
		return []byte{}, err
	}

	return decoded, nil
}
