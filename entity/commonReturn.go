package entity

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/7 22:43
 **/

/*
	通用的类型：包含id和error
*/
type CommonReturn struct{
	Id int `json:"id"`
	Error string `json:"error"`
}