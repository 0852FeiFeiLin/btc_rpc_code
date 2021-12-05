package request

import (
	"btc_RPC_code/util"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
	自己创建一个客户端，然后设置请求头等信息，客户端.Do发起请求
*/

func RPCRequest(method, url string, body io.Reader) (re []byte, err error) { //方法，url和管道流，
	//创建客户端
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	//创建请求
	request, err2 := http.NewRequest(method, url, body)
	if err2 != nil {
		return nil,err
	}

	//base64编码
	/*baseStr := util.Base64Encode(conf.RPC_USER + ":" + conf.RPC_PASSWRD)
	var base64Str = "Basic " + baseStr
*/

	/*request.Header.Set("Encoding", "UTF-8")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", base64Str)
*/
	//客户端设置请求头，遍历map，然后add
	for key,value := range util.Header() {
		request.Header.Add(key,value)
	}

	//客户端发起请求
	resp, err2 := client.Do(request)
	//判断响应码，是否成功
	if resp.StatusCode != 200 {
		//将请求状态码返回
		return nil,errors.New( strconv.Itoa(resp.StatusCode))
	}
	//判断错误信息
	if err != nil {
		return nil,err
	}

	//接收响应结果
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	//返回响应体数据，还没序列化后的
	return respBytes, nil



	////发送请求
	//resp, err := http.Post(RPCURL, "application/json", bytes.NewBuffer(jsonData))//url，json格式,新建管道(放我们的Jsondatade []byte数据，)
	////http.Post("URL","请求数据的资源类型"，"数据的比特流byte[]")
	////数据怎么放到io.reader的管道里面，然后发送数据，新建管道，里面就是我们的
}
