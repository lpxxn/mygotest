package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func EncryptOAEP(secretMessage string, pubkey *rsa.PublicKey) (string, error) {
	//label := []byte("OAEP Encrypted")
	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, pubkey, []byte(secretMessage), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptOAEP(cipherText string, privKey *rsa.PrivateKey) (string, error) {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	//label := []byte("OAEP Encrypted")

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privKey, ct, nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))

	return string(plaintext), nil
}

func SignPKCS1v15(plaintext string, privKey *rsa.PrivateKey) (string, error) {
	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	hashed := sha256.Sum256([]byte(plaintext))
	signature, err := rsa.SignPKCS1v15(rng, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func VerifyPKCS1v15(signature string, plaintext string, pubkey *rsa.PublicKey) (bool, error) {
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	hashed := sha256.Sum256([]byte(plaintext))
	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		if errors.Is(err, rsa.ErrVerification) {
			return false, nil
		}
		return false, err
	}
	return true, err
}

// 加密
func RsaEncrypt(origData string, pubkey *rsa.PublicKey) (string, error) {
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, []byte(origData))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// 解密
func RsaDecrypt(cipherText string, privKey *rsa.PrivateKey) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	revData, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, data)
	if err != nil {
		return "", err
	}
	return string(revData), nil
}
func ParsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	//var block *pem.Block
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("empty private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ParsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("empty public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}
