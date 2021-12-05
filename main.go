package main

import (
	"btc_RPC_code/router"
	"fmt"
	"net/http"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/15 22:10
 **/
func main() {
	//设置静态资源
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
	router.RouterHandle()

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("server running on 127.0.0.1:8091")
}