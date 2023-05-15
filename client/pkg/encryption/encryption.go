package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"io"
)

func Encode(data any, masterkey string) (resp []byte, err error) {

	key, err := getHexKey(masterkey)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(data)
	if err != nil {
		return nil, err
	}

	plaintext := buf.Bytes()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func Decrypt(encryptData []byte, masterKey string) (resp []byte, err error) {
	key, err := getHexKey(masterKey)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := encryptData[:nonceSize], encryptData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil

}

func Decode[T any](encryptData []byte, masterKey string) (resp any, err error) {
	var t T
	dec, err := Decrypt(encryptData, masterKey)

	test := gob.NewDecoder(bytes.NewReader(dec))
	err = test.Decode(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func getHexKey(key string) ([]byte, error) {
	h := sha256.New()
	h.Write([]byte(key))
	hexKey := hex.EncodeToString(h.Sum(nil))
	return hex.DecodeString(hexKey)
}
