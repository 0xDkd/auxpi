package bootstrap

import (
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
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

func FormatSoftLink(url *string) {
	n := len(*url)
	rs := []rune(*url)
	s := string(rs[n-1:n])
	if s == "/" {
		*url = string(rs[0:n-1])
	}
	s = string(rs[0:1])
	if s != "/" {
		*url = "/"+string(rs[0:n])
	}
}

func FormatStoreLocation(location *string) {
	n := len(*location)
	rs := []rune(*location)
	s := string(rs[n-1 : n])
	if s != "/" {
		*location += "/"
	}
	s = string(rs[0:1])
	if s == "/" {
		*location = string(rs[1:n])
	}
}

func GetRandomString(l int,str string) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//检查路径并且创建
func CheckPath(path string) {
	//base := bootstrap.GetPath()
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0775)
		if err != nil {
			logs.Alert("Create Images store unsuccessful:", err)
			return
		}
	}
}

//格式化 url
func FormatUrl(url *string) {
	n := len(*url)
	rs := []rune(*url)
	s := string(rs[n-1 : n])
	if s != "/" {
		*url += "/"
	}
	s = string(rs[0:1])
	if s == "/" {
		*url = string(rs[1:n])
	}
}