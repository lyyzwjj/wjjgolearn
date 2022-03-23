package _004crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

//var (
//	initialVector = "1234567890123456"
//	passphrase    = "Impassphrasegood"
//)

type properties struct {
	secret1 string
	vector1 string
}

var envProperties = map[string]properties{
	"un-pro": {
		"v6DhzUGM8G2k3TvVtYN4S7Q6qyAW84X8",
		"scwUhZHYP6ZnLhnTq6sgXtuLT54JUFZ5",
	},
	"pro": {
		"mN98mTjOtV4PkNaD5WnyGCSkRvl5GLcF",
		"1Qu4Bd8CnaxGp6wfCm2oVYSf4Nl2vybb",
	},
	"test": {
		"kVqXbVna6Fg7dAED",
		"Mbfb9cLNh56Fivn6",
	},
}

func TestAesUtil(t *testing.T) {
	if properties, ok := envProperties["test"]; !ok {
		fmt.Println("查无此key")
	} else {
		var plainText = "hello world"
		encryptedData := AESEncrypt(plainText, properties.secret1, properties.vector1)
		encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
		fmt.Println(encryptedString)
		encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
		decryptedText := AESDecrypt(encryptedData, properties.secret1, properties.vector1)
		fmt.Println(string(decryptedText))
	}
}

func AESEncrypt(src string, key string, vector string) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(vector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt []byte, key string, vector string) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	// cipherText := make([]byte, aes.BlockSize+len(crypt))

	//iv is the ciphertext up to the blocksize (16)
	// iv := cipherText[:aes.BlockSize]
	//if _, err = io.ReadFull(rand.Reader, iv); err != nil {
	//	return
	//}

	ecb := cipher.NewCBCDecrypter(block, []byte(vector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
