package entity
/*
	该结构体用于描述，获取最新区块的响应值，里面有三个字段
*/
type BestBlockHash struct {
	CommonReturn
	Result string `json:"result"`
}

