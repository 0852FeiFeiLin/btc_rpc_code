package server

import "fmt"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/18 21:20
 **/
/*

 */
func KVServer2(command string) (interface{},error){
	switch command {
	case "getbestblockhash":
		return  GetBestBlockHash()
	case "getblockcount":
		return GetBlockCount()
	case "getbestblockinfo":
		return  GetBestBlockInfo()
	case "getbalance":
		return  GetBalance()
	case "help":
		return  GetHelp()
	case "getmempoolinfo":
		return GetMempoolInfo()
	case "getdifficulty":
		return GetDifficulty()
	case "getblockheader":
		return GetBlockHeader()
	case "getblockchaininfo":
		return GetBlockChainInfo()
	case "listunspent":
		return  ListunSpent()
	case "getblock":
		return GetBlock()
	default:
		fmt.Println("你输入的命令行有误")
	}
	return nil,nil
}
