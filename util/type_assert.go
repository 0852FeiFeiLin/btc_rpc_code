package util

import (
	"btc_RPC_code/entity"
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/2 21:20
 **/
/*
	类型断言：因为前面调用请求，返回的是interface{} 类型的，最终是要转换为byte[] 类型，进行反序列化的，
		但是我们不知道他是什么类型，所以需要switch case...这样来捕捉一遍，
		判断他到底是什么类型，然后进行相应的类型转换，例如interface --> string ---> byte[]
		最终反序列化成我们看的懂的，供我们所需要

*/
//类型断言
func TypeAssert(info interface{}) interface{} {
	/*
		中间用switch...case 判断一遍所有可能出现的类型，捕捉一遍，然后执行相应的操作
	*/

	switch info.(type) {
	case string:
		return info.(string)
		fmt.Println("this is string")
	case int:
		fmt.Println("this is int")
	case float64:
		return info.(float64)

		fmt.Println("this is float64")
	case []uint8:
		return info.([]uint8)
		/*var result1 entity.BlockChainInfo
		err := json.Unmarshal(result, &result1)
		fmt.Println("我断言转换成功了\n", result1)

		return result1, err
		*/
	case interface{}:
		fmt.Println("我断言成啥啊")

		return info.(string)
	case entity.Block:
		return info.(entity.Block)
	default:
		fmt.Println("unknown type")
	}

	return nil
}

/*
type CommonReturn struct {
	Id int `json:"id"`
	Error string `json:"error"`
}

type GetBlockCount struct {
	CommonReturn
	Result int
}
type GetBlockHash struct {
	CommonReturn
	Result string
}*/

func TypeAssert2(info interface{}) interface{} {
	t := reflect.TypeOf(info)
	fmt.Println("类型是：",t)
	switch t.String() {
	case "int":
		return fmt.Sprintf("%d",info.(int))
	case "string":
		return fmt.Sprintf("%s",info.(string))
	case "float64":
		return fmt.Sprintf("%f",info.(float64))
	case "[]uint8":
		return fmt.Sprintf("%s",info.([]uint8))
	case "interface":
		return info.(string)
	case "bool":
		return fmt.Sprintf("%v",info.(bool))
	case "entity.Block":
		//创建字节变量存储结构体转换数据
		var result bytes.Buffer
		//encoder为实际转换者，创建一个转换者
		encoder := gob.NewEncoder(&result)
		gob.Register(entity.Block{})
		//将结构体转换
		err := encoder.Encode(&info)
		if err != nil {
			fmt.Println("转换失败block")
			fmt.Println(err.Error())
			return err
		}
		//返回字节变量
		return result
	default:
		fmt.Println("unknown type")
	}
	return nil
}
