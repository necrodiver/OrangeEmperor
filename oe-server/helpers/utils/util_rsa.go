package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
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
func RSAEncrypt(orgidata []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKeyOfRSA, orgidata)
}

// RSADecrypt 私钥解密
func RSADecrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKeyOfRSA, ciphertext)
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
