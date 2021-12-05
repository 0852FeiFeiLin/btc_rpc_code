package main

import (
	"btc_RPC_code/server"
	"fmt"
)
/*
	思路：
	1. 开启rpc协议服务，配置信息(user,password,ip…)，准备环境，准备数据
	 准备数据
		确定url == "http://127.0.0.1:8332"
		组装请求数据：json格式（结构体）
		请求方法：post（对应json格式）

	2. 编写程序
	发起请求url（method = post,ip,params,和请求地址，指定格式。。。。）
	3. 接收响应（反序列化，存在结构体里面）
	4. 输出打印

*/
func main() {
	fmt.Println("heoll")
	//401 用户没有通过认证，请求头设置Authorization
	server.GetBestBlockHash()

	server.GetBlockCount()

	server.GetBlock()

/*	server.GetBalance()

	server.GetHelp()
	//后期就能通过用户访问help，然后输入命令行

	server.GetDifficulty()

	server.GetBlockHeader()

	server.GetBlockChainInfo()

	server.GetMempoolInfo()

	server.GetListunSpent()*/
}
