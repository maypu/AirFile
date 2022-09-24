package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

// FileEncrypt 文件加密
func FileEncrypt(fileBuffer bytes.Buffer, filepath string) bool {
	outfile, err := os.Create(filepath)
	if err != nil {
		panic(err)
		return false
	}
	defer outfile.Close()
	length := fileBuffer.Len()
	splitLength := SplitLength(fileBuffer)
	size := length / splitLength
	key := base64.StdEncoding.EncodeToString([]byte(GetConfig("common.name")))
	//循环加密
	for i := 0; i < size; i++ {
		block := fileBuffer.Bytes()[i*splitLength : (i+1)*splitLength]
		outfile.Write(BytesCombine(block, []byte(key)))
	}
	//处理剩余部分
	block := fileBuffer.Bytes()[size*splitLength:]
	outfile.Write(block)
	return true
}

// FileDecrypt 文件解密
func FileDecrypt(filepath string) *bytes.Buffer {
	file, err := os.Open(filepath) //打开文件
	if err != nil {
		return nil
		panic(err)
	}
	defer file.Close()

	fileBuffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuffer, file); err != nil {
		panic(err)
		return nil
	}

	outfile := bytes.NewBuffer(nil)
	length := fileBuffer.Len()
	splitLength := SplitLength(*fileBuffer)
	key := base64.StdEncoding.EncodeToString([]byte(GetConfig("common.name")))
	encLen := len(key) + splitLength
	//循环解密
	for i := 0; i <= (length / encLen); i++ {
		var block []byte
		if i == 0 {
			block = fileBuffer.Bytes()[0:splitLength]
		} else if i == (length / encLen) {
			block = fileBuffer.Bytes()[splitLength+(i-1)*encLen:]
		} else {
			block = fileBuffer.Bytes()[splitLength+(i-1)*encLen : splitLength+(i)*encLen]
		}
		decByte := bytes.Replace(block, []byte(key), []byte(""), 1)
		//decByte := bytes.Split(block, []byte(key))[0]
		outfile.Write(decByte)
	}
	return outfile
}

// SplitLength 通过分割长度，最终分成1-10切片
func SplitLength(fileBuffer bytes.Buffer) int {
	key := base64.StdEncoding.EncodeToString([]byte(GetConfig("common.name")))
	fileBytes := bytes.Replace(fileBuffer.Bytes(), []byte(key), []byte(""), -1)
	length := len(fileBytes)
	splitLen := int(length / 10) //向下取整
	if splitLen == 0 {
		splitLen = length
	}
	return splitLen
}
