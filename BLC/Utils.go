package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
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

// 字节数组反转

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
