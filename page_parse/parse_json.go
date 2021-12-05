package page_parse

import (
	"btc_RPC_code/entity"
	"encoding/json"
	"fmt"
)

/*
	这个函数用来解析数据
*/

func Parse_Json(res []byte)(re entity.BestBlockHash,err error){
	//解析响应结果
	var bestbalockhash entity.BestBlockHash
	err = json.Unmarshal(res, &bestbalockhash)
	if err != nil {
		fmt.Println("反序列化失败")
		fmt.Println(err.Error())
		return
	}

	//打印输出
	//fmt.Println(bestbalockhash)
	//返回反序列化后的数据
	return bestbalockhash,nil
}