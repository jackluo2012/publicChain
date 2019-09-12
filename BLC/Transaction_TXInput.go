package BLC

type TXInput struct {
	//
	TxHash []byte
	// 2.存储的TXOutput 里面的索引
	Vout int
	//用户名
	ScriptSing string
}