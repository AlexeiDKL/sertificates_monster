package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

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
		log.Println("Create PrivateKey Error: ", err)
		return nil, err
	}
	err = WritePrivateKey(privateKey, path)
	if err != nil {
		log.Println("WritePrivateKey Error: ", err)
		return nil, err
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
	log.Println(key)
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
