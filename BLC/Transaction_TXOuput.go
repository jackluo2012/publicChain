package BLC

type TXOutput struct {
	//多少钱
	Value  int64
	//用户名
	ScriptPubKey string
}
/**
 * 判断当前的消费是谁的钱 是否是 一一对应
 */
func (txOutput *TXOutput) UnLockWithAdress(address string) bool {

	return txOutput.ScriptPubKey == address
}
