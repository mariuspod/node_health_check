package node_health_check

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type JsonRpcRequest struct {
	Id      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
}

type JsonRpcSyncingResponse struct {
	Id      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  bool   `json:"result"`
}

type JsonRpcBlockNumberResponse struct {
	Id      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func GetSyncingState(rpcUrl string) bool {
	reqObj := &JsonRpcRequest{
		Id:      1,
		JsonRpc: "2.0",
		Method:  "eth_syncing",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(reqObj)
	req, err := http.NewRequest("POST", rpcUrl, payloadBuf)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var resObj JsonRpcSyncingResponse
	json.Unmarshal(body, &resObj)
	if err != nil {
		log.Fatal(err.Error())
	}
	return resObj.Result
}

func GetBlockHeight(rpcUrl string) int {
	reqObj := &JsonRpcRequest{
		Id:      1,
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(reqObj)
	req, err := http.NewRequest("POST", rpcUrl, payloadBuf)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var resObj JsonRpcBlockNumberResponse
	json.Unmarshal(body, &resObj)
	if err != nil {
		log.Fatal(err.Error())
	}

	block, err := hexutil.DecodeUint64(resObj.Result)
	if err != nil {
		log.Fatal(err.Error())
	}
	return int(block)
}
