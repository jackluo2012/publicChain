package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

//标准的JSON字符串转数组

func JSONToArray(jsonstring string) []string {

	//json 到 []string

	var sArr []string

	if err := json.Unmarshal([]byte(jsonstring), &sArr); err != nil {
		log.Panic(err)
	}
	return sArr

}
