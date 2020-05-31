package utils

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

var (
	publicKeyOfRSA  *rsa.PublicKey
	privateKeyOfRSA *rsa.PrivateKey
)

func init() {
	rsaCreateKey(256)
}

// rsaCreateKey RSA公钥私钥产生
func rsaCreateKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// _" + time.Now().Format("2006-01-02_15-04-05") + "
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
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
	// _" + time.Now().Format("2006-01-02_15-04-05") + "
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	publicKeyOfRSA = publicKey
	privateKeyOfRSA = privateKey
	return nil
}

// RSAEncrypt RSA公钥加密
func RSAEncrypt(orgidata []byte) (string, error) {
	// encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyOfRSA, orgidata)
	// if err != nil {
	// 	return "", err
	// }
	// return base64.URLEncoding.EncodeToString(encryptedStr), nil

	partLen := publicKeyOfRSA.N.BitLen()/8 - 11
	chunks := split(orgidata, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyOfRSA, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}
	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

// RSADecrypt 私钥解密
func RSADecrypt(ciphertext []byte) (string, error) {
	// decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyOfRSA, ciphertext)
	// if err != nil {
	// 	return "", err
	// }
	// return string(decrypted), nil

	partLen := privateKeyOfRSA.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(string(ciphertext))
	chunks := split([]byte(raw), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyOfRSA, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	return buffer.String(), err
}

// RsaSignWithSha256 RSA签名
func RsaSignWithSha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKeyOfRSA, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}
	return signature
}

//RsaVerySignWithSha256 RSA验签
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKeyOfRSA, crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
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
