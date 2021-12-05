package conf

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/10/27 11:41
 **/

/*
	关于比特币客户端rpc功能调用的相关配置信息，url、user、等都和Bitcoin.conf中相关配置相关配置选项保持一致
*/
const RPCURL = "http://127.0.0.1:8332"
const RPC_USER = "feifei"
const RPC_PASSWRD = "085200"

/*
	当前程序已完成的可以使用的rpc功能
*/
const (
	GETBESTBLOCKHASH  = "getbestblockhash"
	GETBLOCKCOUNT     = "getblockcount"
	GETBLOCK          = "getblock"
	GETBALANCE        = "getbalance"
	HTLP              = "help"
	GETMEMPOOLINFO    = "getmempoolinfo"
	GETDIFFICULTY     = "getdifficulty"
	GETBLOCKHEADER    = "getblockheader"
	GETBLOCKCHAININFO = "getblockchaininfo"
	LISTUNSPENT       = "listunspent"
)
