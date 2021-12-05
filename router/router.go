package router

import (
	"btc_RPC_code/active"
	"net/http"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/15 21:54
 **/

func RouterHandle(){
	http.HandleFunc("/",active.ToIndex)
	http.HandleFunc("/index",active.Index)

}
