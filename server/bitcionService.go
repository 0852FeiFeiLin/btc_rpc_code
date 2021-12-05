package server

import (
	"btc_RPC_code/conf"
	"btc_RPC_code/page_parse"
	"btc_RPC_code/request"
	"btc_RPC_code/util"
	"bytes"
	"encoding/json"
	"fmt"
)

//获取bestblockhash，并返回
func GetBestBlockHash() (string, error) {
	/*	reqData := entity.RPCReqData{
		//	Id:      time.Now().Unix(),
		Id:      10001,
		Method:  "getbestblockhash",
		Params:  nil,
		JsonRPC: "2.0",
	}*/

	//[]interface{}  就算传递nil他也把他当成一个对象。
	//... param    传参的时候可以不传递参数
	reqData := util.BTC_Prepare(10001, conf.GETBESTBLOCKHASH)

	//序列化结构体数据
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}
	//fmt.Println("组装好的json数据：", string(jsonData))

	//发请求方法,获得响应
	rpcresponse, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonData)) //参数三是创建一个管道
	if err != nil {
		return "", err
	}

	//解析响应数据
	bestblockhash, err := page_parse.Parse_Json(rpcresponse)

	//输出bestblockhash
	fmt.Println("bestblockhash：", bestblockhash.Result)

	return bestblockhash.Result, nil
}

/*
	该函数用于调用比特币的rpc的getblockcount功能，并将获取的当前节点区块总数返回
	如果在请求过程中遇到错误，则返回-1
*/
func GetBlockCount() (interface{}, error) {
	//准备数据
	reqData := util.BTC_Prepare(10001, conf.GETBLOCKCOUNT)
	//序列化结构体数据
	jsonReq, err := json.Marshal(reqData)
	if err != nil {
		return -1, err
	}
	//调用请求
	rpcresp, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	//解析
	blockcount, err := page_parse.Parse_Json1(rpcresp)
	//输出hash总数
	fmt.Println("blockcount:", blockcount.Result)

	return blockcount.Result, nil
}

//获取最新的区块信息
func GetBestBlockInfo() (interface{}, error) {
	//先获取最新区块，或者后期用户输入区块hash然后传递
	besthash, _ := GetBestBlockHash()
	//准备数据
	reqtData := util.BTC_Prepare(10002, conf.GETBLOCK, besthash)
	//序列化结构体数据
	jsonRpc, err := json.Marshal(reqtData)
	if err != nil {
		return nil, err
	}
	//调用请求
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonRpc))
	//解析
	/*blockData, err := page_parse.Parse_Json1(res)
	if err != nil {
		return nil, err
	}
	//输出打印
	fmt.Println("最新区块的信息详情：", blockData.Result)
	return blockData.Result, nil*/

	//换成了自己写的那个结构体
	/*var block entity.Block
	fmt.Println(string(res))
	err = json.Unmarshal(res, &block)
	*/

	if err != nil {
		return nil, err
	}
	return res,nil
}

//查看当前比特币客户端余额
func GetBalance() (interface{}, error) {
	reqData := util.BTC_Prepare(10001, conf.GETBALANCE)
	jsonreq, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonreq))
	blance, err := page_parse.Parse_Json1(res)
	//util.TypeAssert(blance.Result)
	fmt.Println(string(res))
	fmt.Println("钱包余额：", blance.Result)
	return blance.Result, nil

}

/*获取help命令*/
func GetHelp() (interface{}, error) {
	reqData := util.BTC_Prepare(10001, conf.HTLP)
	jsonReq, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	//fmt.Println(string(res))
	if err != nil {
		return nil, err

	}
	help, err := page_parse.Parse_Json1(res)
	fmt.Println("help命令集:\n", help.Result)
	return help.Result, err
}

//获取出块难度
func GetDifficulty() (interface{}, error) {
	//准备数据
	reqData := util.BTC_Prepare(10001, conf.GETDIFFICULTY)
	//数据序列化
	jsonReq, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	//发起请求，返回响应
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	//解析响应数据
	difficult, err := page_parse.Parse_Json1(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//返回响应数据
	return difficult.Result, nil
}

//获取区块链当前信息
func GetBlockChainInfo() (interface{}, error) {
	reqData := util.BTC_Prepare(10001, conf.GETBLOCKCHAININFO)
	jsonReq, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	fmt.Println(string(res))
	blockChainInfo, err := page_parse.Parse_Json1(res)
	fmt.Println(blockChainInfo.Result)
	return res, nil
}

//获取指定区块头
func GetBlockHeader() (interface{}, error) {
	hash, _ := GetBestBlockHash()
	reqData := util.BTC_Prepare(10001, conf.GETBLOCKHEADER, hash)
	jsonReq, _ := json.Marshal(reqData)
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	blockHeader, err := page_parse.Parse_Json1(res)
	return blockHeader.Result, nil
}

//获取区块池信息
func GetMempoolInfo() (interface{}, error) {
	reqData := util.BTC_Prepare(10001, conf.GETMEMPOOLINFO)
	jsonReq, _ := json.Marshal(reqData)
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	mempoolInfo, err := page_parse.Parse_Json1(res)
	if err != nil {
		return nil, err
	}
	return mempoolInfo.Result, nil
}

//获取当前节点的所有交易输出
func ListunSpent() (interface{}, error) {
	reqData := util.BTC_Prepare(10001, conf.LISTUNSPENT)
	jsonReq, _ := json.Marshal(reqData)
	res, err := request.RPCRequest("POST", conf.RPCURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	listtunspent, err := page_parse.Parse_Json1(res)
	if err != nil {
		return nil, err
	}
	return listtunspent.Result, nil

}
/*
	该函数用于获取具体区块的详细信息，并返回，接收一个参数,然后添加参数发起请求
*/
func GetBlock(hash string)(interface{},error){
	return nil,nil
}