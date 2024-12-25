package util

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
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

// 压缩字符串
func CompressString(input string) ([]byte, error) {
	var buf bytes.Buffer
	zw, err := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = zw.Write([]byte(input))
	if err != nil {
		return nil, err
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecompressString 解压缩字符串
func DecompressString(compressed []byte) (string, error) {
	b := bytes.NewReader(compressed)
	zr, err := zlib.NewReader(b)
	if err != nil {
		return "", err
	}
	defer zr.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型(占4个字节)
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头 小端
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) ([]byte, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return nil, err
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return nil, err
	}
	return pack[4:], nil
}
