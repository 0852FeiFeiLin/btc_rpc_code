package entity
/*
	这能不能当成返回的结果集结构体
*/
type Result struct {
	Id int `json:"id"`
	Error string `json:"error"`
	//返回的结果集字段
	Result interface{} `json:"result"`

}
