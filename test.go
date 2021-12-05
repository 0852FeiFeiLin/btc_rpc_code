package main

import "btc_RPC_code/server"

/*
	这是一个测试文件
*/
func main() {
	/*difficulty, _ := server.GetDifficulty()
	fmt.Println(difficulty)
	info, _ := server.GetBlockChainInfo()
	fmt.Println(info)
	header,err:= server.GetBlockHeader()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(header)*/
	/*info, err:= server.GetMempoolInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)*/
/*	spent, err := server.GetListunSpent()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(spent)*/
	//server.GetBalance()
	//server.GetBlockChainInfo()


	server.GetBlockChainInfo()

}
