package management

import (
	"fmt"

	"dkl.dklsa.certificates_monster/iternal/certificate"
	"dkl.dklsa.certificates_monster/iternal/cryptoKey"
	storage "dkl.dklsa.certificates_monster/iternal/storage/mssql"
)

func Certifcate(path string) (string, error) {
	// получаем publik key
	// получаем фразу
	// создаем сертификат
	// сохраняем сертификат в базу
	// возвращаем сертификат
	privateKey, err := cryptoKey.GetPrivateKey(path)
	if err != nil {
		return "",
			fmt.Errorf("Error creating private key: %s", err.Error())
	}

	publicKey := cryptoKey.CreatePublicKeys(*privateKey)
	phrase, err := storage.GetPhrase()
	if err != nil {
		return "", fmt.Errorf("Error getting phrase: %s", err.Error())
	}

	certificate, err := certificate.CreateCertificate(phrase, &publicKey)
	if err != nil {
		return "",
			fmt.Errorf("Error creating certificate: %s", err.Error())
	}
	err = storage.SaveCertificate(phrase, certificate)
	if err != nil {
		return "",
			fmt.Errorf("Error saving certificate: %s", err.Error())
	}
	fmt.Println("Certificate created and saved!")
	return certificate, nil
}
