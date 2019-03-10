package utils

import (
	"auxpi/bootstrap"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetSha256CodeWithSalt(s string) string {
	h := sha256.New()
	salt := bootstrap.SiteConfig.AuxpiSalt
	h.Write([]byte(s + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetMd5CodeWithSalt(s string) string {
	h := md5.New()
	WTFTime := time.Now().Unix()
	SWTFTime := strconv.Itoa(int(WTFTime))
	h.Write([]byte(s + SWTFTime + bootstrap.SiteConfig.AuxpiSalt))
	bs := h.Sum(nil)
	beego.Alert(hex.EncodeToString(bs), WTFTime)
	return hex.EncodeToString(bs)
}

const (
	CHAR_SET               = "UTF-8"
	BASE_64_FORMAT         = "UrlSafeNoPadding"
	RSA_ALGORITHM_KEY_TYPE = "PKCS8"
	RSA_ALGORITHM_SIGN     = crypto.SHA256
)

type XRsa struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// 生成密钥对
func CreateKeys(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}
	derStream := MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}

	return nil
}

func NewXRsa() (*XRsa, error) {
	//自动检测并且生成秘钥
	publicKey, privateKey, err := autoCreate()
	if err != nil {
		return nil, err
	}

	//然后获取
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	block, _ = pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri, ok := priv.(*rsa.PrivateKey)
	if ok {
		return &XRsa{
			publicKey:  pub,
			privateKey: pri,
		}, nil
	} else {
		return nil, errors.New("private key not supported")
	}
}

// 公钥加密
func (r *XRsa) PublicEncrypt(data string) (string, error) {
	logs.Alert(r.publicKey)
	partLen := r.publicKey.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}

	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

// 私钥解密
func (r *XRsa) PrivateDecrypt(encrypted string) (string, error) {
	partLen := r.publicKey.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(encrypted)
	chunks := split([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}

	return buffer.String(), err
}

// 数据加签
func (r *XRsa) Sign(data string) (string, error) {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, RSA_ALGORITHM_SIGN, hashed)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(sign), err
}

// 数据验签
func (r *XRsa) Verify(data string, sign string) error {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	decodedSign, err := base64.RawURLEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(r.publicKey, RSA_ALGORITHM_SIGN, hashed, decodedSign)
}

func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)

	k, _ := asn1.Marshal(info)
	return k
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func autoCreate() ([]byte, []byte, error) {
	_, pv := os.Stat("pem/private_key.pem")
	_, pb := os.Stat("pem/public_key.pem")
	if pv != nil && pb != nil {
		bootstrap.CheckPath("pem/")
		privateKey, _ := os.OpenFile("pem/private_key.pem", os.O_CREATE|os.O_WRONLY, 0755)
		publicKey, _ := os.OpenFile("pem/public_key.pem", os.O_CREATE|os.O_WRONLY, 0755)
		err := CreateKeys(publicKey, privateKey, 2048)

		if err != nil {
			logs.Alert(err)
			return []byte{}, []byte{}, err
		}
	}
	pvk, _ := os.OpenFile("pem/private_key.pem", os.O_RDONLY, 0755)
	pbk, _ := os.OpenFile("pem/public_key.pem", os.O_RDONLY, 0755)

	pvkb, _ := ioutil.ReadAll(pvk)
	pbkb, _ := ioutil.ReadAll(pbk)

	return pbkb, pvkb, nil

}

func GetRsaKey() ([]byte, []byte) {
	pb, pv, _ := autoCreate()
	return pb, pv
}


