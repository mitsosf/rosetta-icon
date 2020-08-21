package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
)

func GetLatestBlock() (string, int64) {

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

	var res interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	switch reflect.TypeOf(res).Kind() {
	case reflect.Map:
		tmp := reflect.ValueOf(res)
		for key, results := range tmp {
			if key == "result" {
				for key, value := range results {

				}
			}
		}
	}
	//for key, results := range res {
	//	if key == "result" {
	//		for key, value := range results {
	//
	//		}
	//	}
	//}
	////timestamp, _ := strconv.Atoi(res["result"]["time_stamp"])
	//hash := res["result"]["block_hash"]
	//return hash, int64(timestamp)
	wow.String()
	return "awda", 1231
}
