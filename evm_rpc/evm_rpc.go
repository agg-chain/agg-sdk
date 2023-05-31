package main

import (
	"agg-sdk/evm_rpc/model"
	"agg-sdk/sdk/evmsdk"
	"agg-sdk/sdkconfig"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"log"
	"net/http"
)

var (
	evmRpcPort = 8080
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" && r.Method != "OPTIONS" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json

		// 读取以太坊JSON-RPC请求
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// 转发请求到另一个HTTP请求地址
		request := &model.Request{}
		err = json.Unmarshal(body, request)
		if err != nil {
			return
		}

		var resultHexForEvm interface{}
		switch request.Method {
		case "net_version":
			resultHexForEvm = sdkconfig.ChainId.String()
			break
		case "eth_sendRawTransaction":
			broadcastResponse, err := evmsdk.BroadcastTx(request.Params[0].(string))
			if broadcastResponse.Result.DeliverTx.Data == nil {
				break
			}
			decoded, err := base64.StdEncoding.DecodeString(broadcastResponse.Result.DeliverTx.Data.(string))
			if err != nil {
				return
			}
			resultHexForEvm = hexutil.Encode(decoded)
			break
		case "eth_getTransactionCount":
			nocne, err := evmsdk.GetNonce(request.Params[0].(string))
			if err != nil {
				return
			}
			resultHexForEvm = fmt.Sprintf("0x%X", nocne)
			break
		case "eth_getBalance":
			balance, err := evmsdk.GetBalance(request.Params[0].(string))
			if err != nil {
				return
			}
			resultHexForEvm = fmt.Sprintf("0x%X", balance)
			break
		case "agg_getTx":
			txs, err := evmsdk.GetTx(request.Params[0].(string), request.Params[1].(string), request.Params[2].(string))
			if err != nil {
				return
			}
			var resultTx []*model.TxDetailsInfo
			err = json.Unmarshal([]byte(txs), &resultTx)
			if err != nil {
				return
			}
			resultHexForEvm = resultTx
			break
		}

		response := model.Response{
			Jsonrpc: request.Jsonrpc,
			Id:      request.Id,
			Result:  resultHexForEvm,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	})

	fmt.Printf("Listen port: %v\n", evmRpcPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", evmRpcPort), nil))

}
