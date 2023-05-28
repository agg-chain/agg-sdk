package evmsdk

import (
	"agg-sdk/model"
	"agg-sdk/sdkconfig"
	"agg-sdk/util"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func abciQuery(data string) ([]byte, error) {
	resultStr := util.Get(fmt.Sprintf("%s/abci_query?data=\"%s\"", sdkconfig.RpcUrl, data))

	var result model.QueryResponse
	err := json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(result.Result.Response.Value)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func GetAddressHex(privateKey *ecdsa.PrivateKey) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress.Hex()
}
