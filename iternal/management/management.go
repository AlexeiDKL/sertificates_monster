package management

import (
	"fmt"

	"dkl.dklsa.certificates_monster/iternal/certificate"
	"dkl.dklsa.certificates_monster/iternal/key"
)

func certifcate(path string) {
	// получаем publik key
	// создаем сертификат
	// сохраняем сертификат в базу
	// возвращаем сертификат
	privateKey, err := key.GetPrivateKey(path)
	if err != nil {
		fmt.Println("Error creating private key:", err.Error())
		return
	}
	fmt.Println("Private key created!")
	publicKey := key.CreatePublicKeys(*privateKey)

	certificate, err := certificate.CreateCertificate(&publicKey)
	if err != nil {
		fmt.Println("Error creating certificate:", err.Error())
		return
	}
	err = database.SaveCertificate(certificate)
	if err != nil {
		fmt.Println("Error saving certificate:", err.Error())
		return
	}
	fmt.Println("Certificate created and saved!")
	return
}
