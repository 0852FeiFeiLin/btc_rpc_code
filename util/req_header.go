package util

import (
	"btc_RPC_code/conf"
)

/*
	这里写请求头的一些信息，然后要设置请求头的时候就遍历这个函数里面的数据，add添加key和value进去就好
*/
func Header() map[string]string {
	baseStr := Base64Encode(conf.RPC_USER + ":" + conf.RPC_PASSWRD)
	var base64Str = "Basic " + baseStr

	header := make(map[string]string)
	header["Encoding"]="UTF-8"
	header["Content-Type"]="application/json"
	header["Authorization"]= base64Str
	//返回map
	return header;
}
