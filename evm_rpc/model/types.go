package model

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type QueryResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Response struct {
			Code      int         `json:"code"`
			Log       string      `json:"log"`
			Info      string      `json:"info"`
			Index     string      `json:"index"`
			Key       string      `json:"key"`
			Value     string      `json:"value"`
			ProofOps  interface{} `json:"proofOps"`
			Height    string      `json:"height"`
			Codespace string      `json:"codespace"`
		} `json:"response"`
	} `json:"result"`
}

type BroadcastTxCommitResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		CheckTx struct {
			Code         int           `json:"code"`
			Data         interface{}   `json:"data"`
			Log          string        `json:"log"`
			Info         string        `json:"info"`
			GasWanted    string        `json:"gas_wanted"`
			GasUsed      string        `json:"gas_used"`
			Events       []interface{} `json:"events"`
			Codespace    string        `json:"codespace"`
			Sender       string        `json:"sender"`
			Priority     string        `json:"priority"`
			MempoolError string        `json:"mempoolError"`
		} `json:"check_tx"`
		DeliverTx struct {
			Code      int         `json:"code"`
			Data      interface{} `json:"data"`
			Log       string      `json:"log"`
			Info      string      `json:"info"`
			GasWanted string      `json:"gas_wanted"`
			GasUsed   string      `json:"gas_used"`
			Events    []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
					Index bool   `json:"index"`
				} `json:"attributes"`
			} `json:"events"`
			Codespace string `json:"codespace"`
		} `json:"deliver_tx"`
		Hash   string `json:"hash"`
		Height string `json:"height"`
	} `json:"result"`
}

type TxDetailsInfo struct {
	Id       int64  `json:"id,omitempty"`
	TxHash   string `json:"tx_hash,omitempty"`
	TxFrom   string `json:"tx_from,omitempty"`
	TxTo     string `json:"tx_to,omitempty"`
	RawData  string `json:"raw_data,omitempty"`
	TxValue  string `json:"tx_value,omitempty"`
	TxHeight int32  `json:"tx_height,omitempty"`
	Gas      *Gas   `json:"gas,omitempty"`
}

type Gas struct {
	GasPrice string `json:"gas_price,omitempty"`
	GasLimit uint64 `json:"gas_limit,omitempty"`
	GasUsed  uint64 `json:"gas_used,omitempty"`
}
