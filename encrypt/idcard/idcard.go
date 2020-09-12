package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io"
	"fmt"
	"crypto/rand"
	"strconv"
	"time"
)

// AES-GCM should be used because the operation is an authenticated encryption
// algorithm designed to provide both data authenticity (integrity) as well as
// confidentiality.

// Merged into Golang in https://go-review.googlesource.com/#/c/18803/

// 随机IV的长度
const NONCE_LENGTH = 12

func ExampleNewGCMEncrypter(key, plaintext string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	keyBytes := []byte(key)
	plaintextBytes := []byte(plaintext)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, NONCE_LENGTH)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertextBytes := aesgcm.Seal(nil, nonce, plaintextBytes, nil)

	return hex.EncodeToString(append(nonce, ciphertextBytes...))
}

func ExampleNewGCMDecrypter(key, ciphertext string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	keyBytes := []byte(key)
	nonce, _ := hex.DecodeString(ciphertext[:NONCE_LENGTH*2])
	ciphertextBytes, _ := hex.DecodeString(ciphertext[NONCE_LENGTH*2:])

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintextBytes, err := aesgcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintextBytes)
}

// 随机生成18位身份证号
func RandIdCard() string {
	ret := ""
	cardBytes := make([]byte, 18)
	if _, err := io.ReadFull(rand.Reader, cardBytes); err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(cardBytes); i++ {
		ret = ret + strconv.Itoa(int(cardBytes[i])%10)
	}
	return ret
}

func main() {
	// 1. 测试AES256-gcm算法
	key := "AES256Key-32Characters1234567890"
	idCarad := RandIdCard()
	fmt.Printf("idCard: %s\n", idCarad)

	cipherText := ExampleNewGCMEncrypter(key, idCarad)
	fmt.Printf("cipherText: %s\n", cipherText) //12字节随机IV，18字节密文，16字节mac校验值

	plainText := ExampleNewGCMDecrypter(key, cipherText)
	fmt.Printf("plainText: %s\n", plainText)

	// 2. 测试AES-gcm加解密时间性能
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		cipherText = ExampleNewGCMEncrypter(key, idCarad)
	}
	elapsedEncryption := time.Since(t1)
	fmt.Println("Encryption elapsed: ", elapsedEncryption)

	t2 := time.Now()
	for i := 0; i < 1000000; i++ {
		ExampleNewGCMDecrypter(key, cipherText)
	}
	elapsedDecryption := time.Since(t2)
	fmt.Println("Decryption elapsed: ", elapsedDecryption)
}
