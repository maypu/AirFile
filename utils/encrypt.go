package utils

import (
	"bytes"
	"io"
	"os"
)

func Encrypter(fileBuffer bytes.Buffer, filepath string) bool {
	outfile, err := os.Create(filepath)
	if err != nil {
		panic(err)
		return false
	}
	defer outfile.Close()
	length := fileBuffer.Len()
	size := length / 245
	//循环加密
	for i := 0; i < size; i++ {
		block := fileBuffer.Bytes()[i*245 : (i+1)*245]
		outfile.Write(BytesCombine(block, []byte(GetConfig("common.name"))))
	}
	//处理剩余部分
	block := fileBuffer.Bytes()[size*245:]
	outfile.Write(BytesCombine(block, []byte(GetConfig("common.name"))))
	return true
}

func Decrypter(filepath string) *bytes.Buffer {
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
	encLen := 245 + len(GetConfig("common.name"))
	//循环解密
	for i := 0; i < (length / encLen); i++ {
		block := fileBuffer.Bytes()[i*encLen : (i+1)*encLen]
		outfile.Write(bytes.Split(block, []byte(GetConfig("common.name")))[0])
	}
	return outfile
}
