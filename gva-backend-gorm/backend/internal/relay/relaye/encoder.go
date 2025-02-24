package relaye

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type Encoder struct {
	Key string
}

func NewEncoder(key string) *Encoder {
	return &Encoder{Key: padKey(key)}
}

func (e *Encoder) Encode(data map[string]any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling data: %w", err)
	}

	encryptedData := e.AES256Encrypt(jsonData)
	return base64.URLEncoding.EncodeToString(encryptedData), nil
}

func (e *Encoder) Decode(encodedData string) (map[string]any, error) {
	decodedData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, fmt.Errorf("error decoding data: %w", err)
	}

	decryptedData := e.AES256Decrypt(decodedData)
	var result map[string]any
	err = json.Unmarshal(decryptedData, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %w", err)
	}

	return result, nil
}

func (e *Encoder) AES256Encrypt(plainText []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic(err.Error()) // This is just for demonstration; handle properly in real code
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error()) // This is just for demonstration; handle properly in real code
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return cipherText
}

func (e *Encoder) AES256Decrypt(cipherText []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic(err.Error()) // This is just for demonstration; handle properly in real code
	}

	if len(cipherText) < aes.BlockSize {
		panic("Ciphertext too small") // This is just for demonstration; handle properly in real code
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText
}

func padKey(key string) string {
	keyBytes := []byte(key)
	padding := 32 - len(keyBytes)%32
	for i := 0; i < padding; i++ {
		keyBytes = append(keyBytes, byte(padding))
	}
	return string(keyBytes)
}
