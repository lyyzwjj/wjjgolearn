package _004crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

func TestRsaUtil(t *testing.T) {
	// key, err := rsa.GenerateKey(rand.Reader, 128)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}

	//生成私钥
	pkcs1PrivateKey := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1PrivateKey,
	}
	//写入文件
	file, err := os.Create("private.pem")
	_ = pem.Encode(file, block)

	//产生公钥 主要取地址
	PublicKey := &key.PublicKey
	//公钥从私钥中产生
	pkixPublicKey, err := x509.MarshalPKIXPublicKey(PublicKey)
	block1 := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pkixPublicKey,
	}
	file2, err := os.Create("public.pem")
	encode := pem.Encode(file2, block1)
	fmt.Println(encode)

}
