package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var (
	CryptKey = []byte("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A")
)

func Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(CryptKey)
	if err != nil {
		return nil, err
	}

	// 使用CTR模式
	iv := make([]byte, aes.BlockSize) // 初始向量
	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 可选：对密文进行Base64编码以便于文本传输
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	return []byte(encodedCiphertext), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	// 如果密文是Base64编码的，先解码
	decodedCiphertext, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(CryptKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(decodedCiphertext))
	stream.XORKeyStream(plaintext, decodedCiphertext)

	return plaintext, nil
}
