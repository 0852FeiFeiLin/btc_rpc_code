package entity
/*
	定义rpc请求的字段，包含是个字段
*/
type RPCReqData struct {
	Id int  `json:"id"`
	Method string `json:"method"`
	Params  []interface{} `json:"params"`
	JsonRPC string `json:"jsonrpc"`
}