package BLC

type TXInput struct {
	//
	TxHash []byte
	// 2.存储的TXOutput 里面的索引
	Vout int
	//用户名
	ScriptSing string
}

/**
 * 判断当前的消费是谁的钱 是否是 一一对应
 */
func (txInput *TXInput) UnLockWithAdress(address string) bool {

	return txInput.ScriptSing == address
}
