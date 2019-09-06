package util

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

func TestOAEP(t *testing.T) {
	// Generate Alice RSA keys Of 2048 Buts
	alicePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Extract Public Key from RSA Private Key
	alicePublicKey := &alicePrivateKey.PublicKey
	secretMessage := "Hello wrod"
	encryptedMessage, err := EncryptOAEP(secretMessage, alicePublicKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Cipher Text  ", encryptedMessage)
	decStr, err := DecryptOAEP(encryptedMessage, alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	if decStr != secretMessage {
		t.Error("not equal")
	}
}

func TestDecryptOAEP(t *testing.T) {
	alicePrivateKey, alicePublicKey, err := ParseKey()
	if err != nil {
		t.Fatal(err)
	}
	secretMessage := "Hello wrod"
	encryptedMessage, err := EncryptOAEP(secretMessage, alicePublicKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Cipher Text  ", encryptedMessage)
	decStr, err := DecryptOAEP(encryptedMessage, alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	if decStr != secretMessage {
		t.Error("not equal")
	}

}

func TestSignPKCS1v15(t *testing.T) {
	// Generate Alice RSA keys Of 2048 Buts
	alicePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Extract Public Key from RSA Private Key
	alicePublicKey := &alicePrivateKey.PublicKey
	secretMessage := "Hello wrod"
	fmt.Println("Original Text  ", secretMessage)
	signature, err := SignPKCS1v15(secretMessage, alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Singature :  ", signature)
	verif, err := VerifyPKCS1v15(signature, secretMessage, alicePublicKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(verif)
}

func TestRsaDecrypt(t *testing.T) {
	// Generate Alice RSA keys Of 2048 Buts
	alicePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Extract Public Key from RSA Private Key
	alicePublicKey := &alicePrivateKey.PublicKey
	secretMessage := "Hello wrod"

	ciphertext, err := RsaEncrypt(secretMessage, alicePublicKey)
	if err != nil {
		t.Error(err)
		return
	}
	revBytes, err := RsaDecrypt(ciphertext, alicePrivateKey)
	if err != nil {
		t.Error(err)
		return
	}
	if revBytes != secretMessage {
		t.Error("not equal")
	}
}

func TestRsaEncrypt2(t *testing.T) {
	alicePrivateKey, alicePublicKey, err := ParseKey()
	if err != nil {
		t.Fatal(err)
	}
	secretMessage := "Hello wrod"

	cipherText, err := RsaEncrypt(secretMessage, alicePublicKey)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cipherText)
	revBytes, err := RsaDecrypt(cipherText, alicePrivateKey)
	if err != nil {
		t.Error(err)
		return
	}
	if revBytes != secretMessage {
		t.Error("not equal")
	}
}

func TestRsaDecrypt2(t *testing.T) {
	encryptedMessage := "ceiFC8JZGqQEgQxm6zp8+e2sBydreSTYT9O5Q0rzj9ByVblQFYmYUpBZoCuGqNE1A/numdv5hBsrlXCc36sbcphncy1dWm/rjJYTfXaws9lqWh9L8CR/7fVQ15KhnBfQyl1MF5gWo3cJyJuulBMjk0DBEsCsN+Jaw4qYS1uAbfkoYYf7eRxEL2PMc1jdUFvYyBbtjTRUEcGYQsvy8/NoBg5digDu/pKRfA86tEYlV4KYLdOnoInPxZIyoRlvXWyi2z4zgTr9oWT9xel8Tvl4J2eJ22/bxy02S+8Nn+1bPE5mVLbwEk960SxVysL8k5cl6WDagN4ThWahZx6drMv0dw=="
	alicePrivateKey, _, err := ParseKey()
	if err != nil {
		t.Fatal(err)
	}

	decStr, err := RsaDecrypt(string(encryptedMessage), alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	secretMessage := "Hello wrod"

	if string(decStr) != secretMessage {
		t.Error("not equal")
	}
}

func ParseKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := ParsePrivateKey([]byte(privateKeyStr))
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	publicKey, err := ParsePublicKey([]byte(publicKeyStr))
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}

// 注意不能有 \n
const privateKeyStr = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxFDJAtX2iWfm3uSPpDyOR6TNAgZjkqP2Vk2b2abRZGUwZdxT
iJ0b01JKMttazGYZbigCDYiXiF5vndcLW7JPc7g02SzrdSUnKVHign36xLzltxI6
OtFp9csIALAxpW9+bWBHoPba5v5RhPl9uERJhcfagZwv5RsG9vnCTarnNeE2D0OC
Sx+BkNGqPfswYLqZ4qrRkzQruuhzJnAGCegr+Bfm2+FMscgWlZ/lGn8V0iDrJJOm
aYx7QHGIqpfkNtsTRD08Aos+ZJkpVIcgJHC5BA4Gvpj5yRwgZaRWDPrH/rgKNxBC
h3XERplzaCaMLqWmXEBNRqcH8+T6sLr7iQvbTwIDAQABAoIBAGCtL/BXkgtkwYsu
n5ZPw/mALP08TDjgyeUCXye2QRPhyQ3AjNqQFC3jLmsRHwvdk9Rtm2UyyPz9FJZ0
WkGT9BCWvF/fAS01WK2+h3+frqjIJZAcR9TDYB0NGSnNmwVPfME8UsnwBykOQPLt
Mn5WEqbxHqgJVITM5IcnxAofibE8E6pCrld+OmyqB+RxLEpE8nAueK6teSw/jdtM
Y50fCNc0uN/EVan+S5COBbAWbc8kVj66y76ntTWMPqLfllpJ9dsdvPuBawRDSMaf
wtaiK6ZEGKBXla6Y8jjt0kRLdSMC9xFWoyfUxfoR7e5iJb2mIMYjV6E313o3CWe7
vJJwB4ECgYEA81TpWzkp8Se2YiTLRHhnEjO2wX66EzGzt8+jgLqNeRcgGVvFw1Q4
2hX7ROLg4ZT9nyLlrw0gBvQUaD31kmOvCfXJ+r5cgWd3wBldDybdqqhiwFFrB/aq
ytE2+WtUfeCpXZHyoNDWTktKvyQY/UOPzPa3mu+L2bWmesvzpZH32W8CgYEAzok/
9MN1di4WmOgcYw1DnNZWPUCDvhCZ0UGuh723VHQQYC01LXY2dsbcTqCCyqZ9mb+U
yTt+dhCr19pRLjoIY3AmH2LwpEpzNtz1at+tqHWHlury6fO0cvYSHhgP1ptR5/Vs
J8cXKFHlEI+hFAEfKHo0PNiHu8xySIF5ziKCbCECgYArwpd86ljan7OzDr7nf2e5
5Eb3oVzBWuzhH5xd6C7NGhWRuelk6V6YCd/4UXzz1KGP3uzDgG7EEV0iKhJh0z+F
YQiD6XwlYYAtseIEeY7fxAyOXZYMBpZMhzeFv4GgaajLxRWHSkdgKTY2Db3Yvi+8
QppeaLoT2xEngEo++nNCTQKBgQCJBySIpvmSnAdgEXdr/EvkjYwP/XKqsmbrH9HE
u4kbnhmTRcsnS2vdKBjYjcPY4dcpH6N+U1uWx37Lqkv8CIKyPqD8Z+9SgRapyrWE
37xerQN1jJH3yDnrn/jrQHmZn9katYi5Z2yk0pjQgQqUtB3RBCsOR3bekUktc8G1
ks3HIQKBgEHlYaGH4C90JCzeprT9MZ0nv1h40h190Ncv+fG6BKGymquHXJhZqTXW
LKhdU6gITftiZVVxOYobj8fBnNHBfxoy9u/duVd3IflSto/IkY347ofSDtPefdy5
dAH2XE7QcmErqgMRDi7TF3LbCxOfR0N6t3edr1kkQjEvneL6+S2P
-----END RSA PRIVATE KEY-----`
const publicKeyStr = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAxFDJAtX2iWfm3uSPpDyOR6TNAgZjkqP2Vk2b2abRZGUwZdxTiJ0b
01JKMttazGYZbigCDYiXiF5vndcLW7JPc7g02SzrdSUnKVHign36xLzltxI6OtFp
9csIALAxpW9+bWBHoPba5v5RhPl9uERJhcfagZwv5RsG9vnCTarnNeE2D0OCSx+B
kNGqPfswYLqZ4qrRkzQruuhzJnAGCegr+Bfm2+FMscgWlZ/lGn8V0iDrJJOmaYx7
QHGIqpfkNtsTRD08Aos+ZJkpVIcgJHC5BA4Gvpj5yRwgZaRWDPrH/rgKNxBCh3XE
RplzaCaMLqWmXEBNRqcH8+T6sLr7iQvbTwIDAQAB
-----END RSA PUBLIC KEY-----`
