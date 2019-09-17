package BLC

type UTXOS struct {
	TxHash []byte
	Index int
	Output *TXOutput
}
