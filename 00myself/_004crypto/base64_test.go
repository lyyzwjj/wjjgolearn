package _004crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64Encrypt(t *testing.T) {
	aliIdSource := "4234c6a478061330"
	aliIdTarget := base64.StdEncoding.EncodeToString([]byte(aliIdSource))
	fmt.Println(aliIdTarget)
	aliSecretSource := "093084dbd30e80ada355808aab04"
	aliSecretTarget := base64.StdEncoding.EncodeToString([]byte(aliSecretSource))
	fmt.Println(aliSecretTarget)

}

func TestBase64Decrypt(t *testing.T) {
	aliIdTarget := "d3g0MjM0YzZhNDc4MDYxMzMw"
	if aliIdSource, err := base64.StdEncoding.DecodeString(aliIdTarget); err == nil {
		fmt.Println(string(aliIdSource))
	}
	aliSecretTarget := "MDkzMDg0ZGJkMzBlODBhZGEzNTU4MDhhYWIwNGNlY2I="
	if aliSecretSource, err := base64.StdEncoding.DecodeString(aliSecretTarget); err == nil {
		fmt.Println(string(aliSecretSource))
	}
}
