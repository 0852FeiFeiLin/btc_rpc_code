package util

import (
	"encoding/base64"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/10/27 10:45
 **/

/*
	该函数用于给定的字符串进行base64编码，返回编码后的结果
*/
func Base64Encode(str string)string{
	return base64.StdEncoding.EncodeToString([]byte(str))
}

/*
该函数用于给定字符串解码，并返回解码结果
*/
func Base64Decode (str string)([]byte,error){
	return  base64.StdEncoding.DecodeString(str)
}
