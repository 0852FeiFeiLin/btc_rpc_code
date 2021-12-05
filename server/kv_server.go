package server

import (
	"fmt"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/18 19:36
 **/
func KVServer(command string) interface{} {
	result := make(map[string]interface{})
	/*
		result["getblock"],_ = server.GetBlock()*/
	result["getbalance"], _ = GetBalance()
	result["getblockcount"] ,_ = GetBlockCount()
	result["getbestblockhash"], _ = GetBestBlockHash()
	//result["help"], _ = GetHelp()
	/*	result["getmempoolinfo"],_ = server.GetMempoolInfo()
		result["getdifficulty"],_ = server.GetDifficulty()
		result["getblockheader"],_ = server.GetBlockHeader()
		result["getblockchaininfo"],_ = server.GetBlockChainInfo()
		result["listunspent"],_ = server.GetListunSpent()*/
	for key, value := range result { //遍历result集合
		if command != key {
			continue
			//fmt.Println(value)
		}
		fmt.Println("-------------------------------------------------------------------------------------------")
		return value
	}
	return nil
}
