package entity

import "math/big"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/7 22:25
 **/

type Block struct {
	//返回通用型
	CommonReturn
	BlockResult `json:"result"`//详细信息
}
/*
	区块的详细信息
*/
type BlockResult struct {
	Bits string `json:"bits"`
	Chainwork string `json:"chainwork"`
	Confirmations int `json:"confirmations"`
	Difficulty int `json:"difficulty"`
	//区块hash值
	Hash string `json:"hash"`
	//区块高度
	Height int `json:"height"`
	//区块的时间戳
	Mediantime big.Int `json:"mediantime"`
	//区块中默克尔树值
	Merkleroot string `json:"merkleroot"`
	//上一块hash
	Previousblockhash string `json:"previousblockhash"`
	//版本
	Version int `json:"version"`
	//区块权重
	Weight int `json:"weight"`

}