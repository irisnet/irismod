package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"math/big"
	"math/rand"
	"time"
)

func RandStr(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strByte := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, strByte[rand.Intn(len(strByte))])
	}
	return result
}

// SymmetricEncrypt 加密
func SymmetricEncrypt(data, key []byte, encryption string) (string, error) {
	//创建加密实例
	var block cipher.Block
	var err error

	if encryption == "aes" {
		block, err = aes.NewCipher(key)
	} else {
		block, err = sm4.NewCipher(key)
	}

	if err != nil {
		return "", err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return hex.EncodeToString(crypted), nil
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// SymmetricDecrypt 解密
func SymmetricDecrypt(data string, key []byte, encryption string) ([]byte, error) {
	dataByte, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return aesDecrypt(dataByte, key, encryption)
}

// AesDecrypt 解密
func aesDecrypt(data, key []byte, encryption string) ([]byte, error) {
	//创建加密实例
	var block cipher.Block
	var err error

	if encryption == "aes" {
		block, err = aes.NewCipher(key)
	} else {
		block, err = sm4.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func GetPrivateFromHex(str string) (*sm2.PrivateKey, error) {
	c := sm2.P256Sm2()
	d, err := hex.DecodeString(str)
	if err != nil {
		return &sm2.PrivateKey{}, err
	}
	k := new(big.Int).SetBytes(d)
	params := c.Params()
	one := new(big.Int).SetInt64(1)
	n := new(big.Int).Sub(params.N, one)
	if k.Cmp(n) >= 0 {
		return &sm2.PrivateKey{}, errors.New("privateKey's D is overflow.")
	}
	privateKey := new(sm2.PrivateKey)
	privateKey.PublicKey.Curve = c
	privateKey.D = k
	privateKey.PublicKey.X, privateKey.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return privateKey, nil
}
