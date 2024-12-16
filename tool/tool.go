package main

import (
	"PortRelay/util"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取命令行参数
	args := os.Args
	if len(args) < 1 {
		panic("args error")
	}

	// help
	if len(args) == 1 || args[1] == "help" {
		Help()
		return
	}

	switch args[1] {
	case "genrsa":
		{
			prvKey, pubKey, _ := GenRsaKey(2048)
			fmt.Println(string(prvKey))
			fmt.Println(string(pubKey))
		}
	case "secret":
		{
			if len(args) < 3 {
				log.Println("args error,must have 2 args")
				return
			}
			fmt.Println(Secret(args[2]))
		}

	}

}

func Help() {
	fmt.Println("Commonand:\n" +
		"	genrsa <bits>   生成rsa密钥对")
}

func Secret(key string) string {
	secretKey, err := util.Encrypt(util.ZeroCopyByte(key))
	if err != nil {
		log.Println("encrypt error:", err)
		return ""
	}
	return string(secretKey)
}

func GenRsaKey(bits int) (prvkey, pubkey []byte, err error) {
	// Generates private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)

	// Generates public key from private key.
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}
