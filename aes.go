package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"log"
)

var aes_sk []byte

//goland:noinspection GoBoolExpressions
func fixSecretKey() {
	if SECRET_KEY == "you must change this value." {
		log.Println("warn: you must change the SECRET_KEY, then recompile")
	}

	if len(SECRET_KEY) < 1 {
		SECRET_KEY = _GetFileSecret(false)
	}

	if len([]byte(SECRET_KEY)) > 32 {
		aes_sk = []byte(SECRET_KEY)[:32]
	} else {
		aes_sk = []byte(SECRET_KEY)
		c := 32 - len([]byte(SECRET_KEY))
		for ; c > 0; c-- {
			aes_sk = append(aes_sk, ' ')
		}
	}
}

func AesEncrypt(data []byte) ([]byte, error) {
	data = bytes.TrimSpace(data)

	buf := bytes.Buffer{}
	buf.Write(data)

	plain, _ := ioutil.ReadAll(&buf)
	for i := aes.BlockSize - len(plain)%aes.BlockSize; i > 0 && i != aes.BlockSize; i-- {
		plain = append(plain, ' ')
	}

	block, err := aes.NewCipher(aes_sk)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plain))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plain)

	return bytes.TrimSpace(ciphertext), nil
}

var _AesDecryptError = errors.New("decryption failed, change the value of `SECRET_KEY` or delete the file `~/.pwgen`")

func AesDecrypt(data []byte) ([]byte, error) {
	data = bytes.TrimSpace(data)

	block, err := aes.NewCipher(aes_sk)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize || len(data)%aes.BlockSize != 0 {
		return nil, _AesDecryptError
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	if len(data) < aes.BlockSize || len(data)%aes.BlockSize != 0 {
		return nil, _AesDecryptError
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	return bytes.TrimSpace(data), nil
}
