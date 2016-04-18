package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var KeyAES []byte

func InitAESKey() (bool, error) {
	home := os.Getenv("HOME")
	if home == "" {
		home = "~/"
	}
	filename := filepath.Join(home, ".docker/aeskey")
	if _, err := os.Stat(filename); err == nil {
		if KeyAES, err = ioutil.ReadFile(filename); err != nil {
			return false, nil
		}
	} else if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// AESEncrypt will encrypt the message with cipher feedback mode with the given key.
func AESEncrypt(message, AESKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(message))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], message)

	return ciphertext, nil
}

// AESDecrypt will decrypt the ciphertext via the given key.
func AESDecrypt(ciphertext, AESKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}
