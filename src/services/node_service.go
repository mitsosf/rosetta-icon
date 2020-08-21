package services

import (
	"bytes"
	"encoding/json"
	"go/types"
	"net/http"
)

type Response struct {
	jsonRpc string
	result  Result
	id      int
}

type Result struct {
	version                  int
	height                   int
	signature                string
	prevBlockHash            string
	merkleTreeRootHash       string
	timeStamp                string
	confirmedTransactionList types.Slice
}

func GetLatestBlock() (string, int64, int64) {

	uri := "https://ctz.solidwallet.io/api/v3"

	values := map[string]string{
		"jsonrpc": "2.0",
		"method":  "icx_getLastBlock",
		"id":      "1234",
	}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	result := res["result"].(map[string]interface{})
	var hash string
	var height int64
	var timestamp int64
	for key, value := range result {
		switch key {
		case "block_hash":
			hash = "0x" + value.(string)

		case "height":
			height = int64(value.(float64))

		case "time_stamp":
			timestamp = int64(value.(float64))
		}
	}
	return hash, height, timestamp
}
