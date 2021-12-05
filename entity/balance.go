package entity

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/7 22:27
 **/

type Balance struct{
	//共同的相同的
	CommonReturn
	Balance string `json:result`
}