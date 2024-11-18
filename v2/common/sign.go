package common

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	KeyTypeHmac    = "HMAC"
	KeyTypeRsa     = "RSA"
	KeyTypeEd25519 = "ED25519"
)


// 获取签名函数，返回(func, error），通过keyType获取函数
// 函数签名：func(string, string) (*string, error)，接受两个string参数，返回一个string指针和error
func SignFunc(keyType string) (func(string, string) (*string, error), error) {
	switch {
	case keyType == KeyTypeHmac:
		return Hmac, nil
	case keyType == KeyTypeRsa:
		return Rsa, nil
	case keyType == KeyTypeEd25519:
		return Ed25519, nil
	default:
		return nil, fmt.Errorf("unsupported keyType=%s", keyType)
	}
}

func Hmac(secretKey string, data string) (*string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey)) // 使用sha256算法和密钥新建一个hmac对象
	_, err := mac.Write([]byte(data))  // 将待加密的数据data转为byte切片，写入hmac对象
	if err != nil {
		return nil, err
	}
	encodeData := fmt.Sprintf("%x", (mac.Sum(nil)))  // mac.Sum(nil)返回hmac签名结果，使用%x转为16进制字符串
	return &encodeData, nil
}

func Rsa(secretKey string, data string) (*string, error) {
	// 解码PEM私钥
	block, _ := pem.Decode([]byte(secretKey))
	if block == nil {
		return nil, errors.New("Rsa pem.Decode failed, invalid pem format secretKey")
	}
	// 解析PKCS#8私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Rsa ParsePKCS8PrivateKey failed, error=%v", err.Error())
	}
	// 类型断言
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Rsa convert PrivateKey failed")
	}
	// 哈希数据，并用私钥签名
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	// 将生成的签名转为base64编码
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return &encodedSignature, nil
}

func Ed25519(secretKey string, data string) (*string, error) {
	block, _ := pem.Decode([]byte(secretKey))
	if block == nil {
		return nil, fmt.Errorf("Ed25519 pem.Decode failed, invalid pem format secretKey")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Ed25519 call ParsePKCS8PrivateKey failed, error=%v", err.Error())
	}
	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Ed25519 convert PrivateKey failed")
	}
	pk := ed25519.PrivateKey(ed25519PrivateKey)
	signature := ed25519.Sign(pk, []byte(data))
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return &encodedSignature, nil
}
