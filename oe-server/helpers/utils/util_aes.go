package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"oe_server/helpers/setting"
)

// AESEntry AES加密,采用GCM方式
func AESEntry(data []byte) ([]byte, error) {
	if setting.AES_KEY == "" {
		panic("AES_KEY获取为空")
	}
	key := []byte(setting.AES_KEY)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, _ := cipher.NewGCM(block)

	// nonce := make([]byte, 12)
	nonce := make([]byte, aesgcm.NonceSize())

	io.ReadFull(rand.Reader, nonce)
	out := aesgcm.Seal(nonce, nonce, data, nil)
	return out, nil
}

// AESDecrypt AES解密
func AESDecrypt(data []byte) ([]byte, error) {
	if setting.AES_KEY == "" {
		panic("AES_KEY获取为空")
	}
	key := []byte(setting.AES_KEY)
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
