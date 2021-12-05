package active

import (
	"btc_RPC_code/server"
	"btc_RPC_code/util"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/15 21:48
 **/
/*
	该函用于get请求，用于跳转到index页面
*/
func ToIndex(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files.Execute(w, nil)

}

/*
	该函数用于post请求，用于index页面的逻辑， 和处理用户输入的命令行，然后显示指定的内容数据再页面
*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello~")
	//创建一个JSON格式的解码对象，用于反序列化，和json.Unmarshar差不多意思
	decoder := json.NewDecoder(r.Body)
	//创建一个map接收反序列化json格式后的数据体
	var parames map[string]string
	//解码body
	decoder.Decode(&parames)
	//通过["参数名称"]
	command := parames["command"]
	fmt.Println("命令行：——----------",command)

	/*err := r.ParseForm()
	if err != nil {
		return
	}
	//用户输入的命令行
	command := r.FormValue("command")*/


	/*
		思路1：封装一个map，通过key 和 value匹配，key是具体命令，value是执行对应的方法，遍历map，看command和key一不一样，
		缺点：他会把每一个value里面的方法执行一遍，影响效率。
		info := server.KVServer(command)
	*/
	/*
		思路2：遍历help命令返回的命令行结果集，然后全部insert进数据库里面，
			然后每次查询数据的command键，看用户输入的命令行正不正确，正确执行对应的操作，
	*/
	/*
		思路3： 那些conf改装成一个map ，key是命令行，value也是具体命令
	*/
	/*
		思路4,：直接通过switch case 解决，判断输入的命令行，然后执行指定的方法
	*/
	info, err := server.KVServer2(command)
	//反射判断类型
	//fmt.Println(reflect.TypeOf(info))
	//files, err := template.ParseFiles("./view/index.html")
	if err != nil {
		return
	}
	//fmt.Println("结果info:",string(info.([]uint8)))
	/*

		初步结果在这里，后面就写好每一个结构体，然后通过下面这种方式，进行类型断言，然后，就一个个对应，、、
		result := info.([]uint8)
		var result1 entity.BlockChainInfo
		err = json.Unmarshal(result, &result1)
	*/
	/*
		//result1, err := util.TypeAssert(info)

		if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result1)*/

/*
	files.Execute(w,info.(string))
	这里就出现问题了，返回的类型是接口类型的，但是具体类型是啥我们不知道，但是最终需要转为string类型，但是直接转string，会报错
*/
	/*判断接口类型*/
/*	typ := reflect.TypeOf(info)
	fmt.Println("我是新方法类型判断的：",typ)*/
	/*
		我这类是类型断言，然后返回还是源类型，然后我后面进行序列化成[]byte 字节数组，后面我在反序列化返回给前端

	ty := util.TypeAssert(info)
	infoBytes, err := json.Marshal(ty)
	if err != nil {
		return
	}
	var ss entity.Result
	err = json.Unmarshal(infoBytes, &ss)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	*/
	info2 := util.TypeAssert2(info)
	//var ss = fmt.Sprintf("%d",tt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//infoBytes, err := json.Marshal(info)
	if info2 == nil {
		fmt.Println("内容为空")
		return
	}
	//后端发送给前端的数据
	io.WriteString(w,info2.(string))

	//files.Execute(w,nil)

	//files.Execute(w,result1)
	return
}
/*
这个也是跳转前端页面的方法， 代码太多了，不采取
func HomePage(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = WriteTemplateToHttpResponse(res, t)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func WriteTemplateToHttpResponse(res http.ResponseWriter,t *template.Template)error{
	if t==nil || res == nil{
		return errors.New("WriteTemplateToHttpResponse: t must not be nil.")
	}
	var buf bytes.Buffer
	err := t.Execute(&buf,nil)
	if err!=nil{
		return err
	}
	res.Header().Set("Content-Type","text/html;charset=utf-8")
	_, err = res.Write(buf.Bytes())
	return err
}
*/