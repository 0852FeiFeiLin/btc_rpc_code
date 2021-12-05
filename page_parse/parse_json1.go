package page_parse

import (
	"btc_RPC_code/entity"
	"encoding/json"
	"fmt"
)
/*
	解析传入的interface{}各种类型
*/
func Parse_Json1(res []byte)(re entity.Result,err error){
	//解析响应结果
	var result entity.Result
	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println("反序列化失败")
		fmt.Println(err.Error())
		return
	}

	//打印输出
	//fmt.Println(result)
	//返回反序列化后的数据
	return result,nil
}
