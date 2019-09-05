package BLC

import (
	"bytes"
	"encoding/binary"
	"github.com/labstack/gommon/log"
)
//将int64转换为64字节数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
