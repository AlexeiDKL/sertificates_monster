package cryptoKey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	// "log"

	"dkl.dklsa.certificates_monster/iternal/file"
)

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	// получаем путь приватного ключа
	// проверяем наличие файла по пути
	// если файла нет, генерируем новый приватный ключ
	// сохраняем приватный ключ в указанный путь
	// возвращаем публичный ключ
	if file.Exists(path) {
		// загружаем приватный ключ из файла
		// возвращаем публичный ключ
		stringPrivateKey := ReadPrivateKey(path)
		return stringPrivateKey, nil
	}
	fmt.Println(path, "not found")
	// генерируем новый приватный ключ и сохраняем его в указанный путь
	privateKey, err := CreatePrivateKey()
	if err != nil {
		return nil, fmt.Errorf("Create PrivateKey Error: %s", err)
	}
	err = WritePrivateKey(privateKey, path)
	if err != nil {
		return nil,
			fmt.Errorf("WritePrivateKey Error: %s", err)
	}
	return privateKey, nil
}

func CreatePrivateKey() (*rsa.PrivateKey, error) {
	rnd := rand.Reader
	key, err := rsa.GenerateKey(rnd, 2048)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func CreatePublicKeys(key rsa.PrivateKey) rsa.PublicKey {
	return key.PublicKey
}

func ReadPrivateKey(path string) *rsa.PrivateKey {
	key := file.GetTextInFile(path)
	// log.Println(key)
	block, _ := pem.Decode([]byte(key))
	keyss, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return keyss
}

func WritePrivateKey(key *rsa.PrivateKey, path string) (err error) {
	pemdataPrivateKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	return file.SaveFile(path, string(pemdataPrivateKey))
}

func EncryptWithPublicKey(text string, publicKey *rsa.PublicKey) (string, error) {
	rng := rand.Reader
	labels := []byte("")
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, publicKey, []byte(text), labels)

	return string(ciphertext), err
}

func DecryptWithPrivateKey(text string, privateKey *rsa.PrivateKey) (string, error) {
	rng := rand.Reader
	labels := []byte("")
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey, []byte(text), labels)

	return string(plaintext), err
}
