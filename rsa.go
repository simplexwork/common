package common

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAEncryptByPubKey 公钥加密
func RSAEncryptByPubKey(key []byte, plaintext []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(crand.Reader, pub, plaintext)
}

// RSADecryptByPriKey 私钥解密
func RSADecryptByPriKey(key []byte, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(crand.Reader, priv, ciphertext)
}

// RSASign 签名
func RSASign(key []byte, src []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	h := crypto.SHA256.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(crand.Reader, priv, crypto.SHA256, hashed)
}

// RSAVerify 验签
func RSAVerify(key []byte, src []byte, sign []byte) error {
	block, _ := pem.Decode(key)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.SHA256.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, sign)
}
