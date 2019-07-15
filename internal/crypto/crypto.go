package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// CreateHash generate sha key
func CreateHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt data
func Encrypt(data []byte, passphrase string) []byte {
	key, _ := hex.DecodeString(passphrase)
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// Decrypt data
func Decrypt(data []byte, passphrase string) ([]byte, error) {
        key, _ := hex.DecodeString(passphrase)
        block, err := aes.NewCipher(key)
        if err != nil {
                panic(err.Error())
        }
        gcm, err := cipher.NewGCM(block)
        if err != nil {
                panic(err.Error())
        }
        nonceSize := gcm.NonceSize()
        nonce, ciphertext := data[:nonceSize], data[nonceSize:]
        plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
        if err != nil {
                panic(err.Error())
        }
        return plaintext, err
}