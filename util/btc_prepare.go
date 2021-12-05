package util

import "btc_RPC_code/entity"

//把数据准备放在一个函数里面，提高代码复用性
func BTC_Prepare(id int,method string,params... interface{}) (entity.RPCReqData){
	//var reqData entity.RPCReqData
	reqData := entity.RPCReqData{
		Id : id,
		Method : method,
		Params : params,
		JsonRPC : "2.0",
	}

	return reqData;
}

