package bootstrap

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func Decode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)

	if err != nil {
		panic(err)
	}
	return string(data)
}

func Encode(enc *base64.Encoding, str string) string {
	bData := []byte(str)
	data := enc.EncodeToString(bData)
	return string(data)
}
