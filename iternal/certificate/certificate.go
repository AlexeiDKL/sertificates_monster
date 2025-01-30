package certificate

import (
	"crypto/rsa"

	"dkl.dklsa.certificates_monster/iternal/cryptoKey"
	"dkl.dklsa.certificates_monster/iternal/storage/mssql"
)

type Certifcate struct {
	Id                int    `json:"id"`
	Phrase            string `json:"phrase"`
	Partner           int    `json:"partner"`
	LicenseStartDate  int    `json:"licenseStart"`
	LicenseEndDate    int    `json:"licenseEnd"`
	LicenseDevicesCNT int    `json:"licenseDevicesCNT"`
	LicenseActivated  string `json:"licenseActivated"`
}

func CreateCertificate(phrase string, publicKey *rsa.PublicKey) (string, error) {
	result, err := cryptoKey.EncryptWithPublicKey(phrase, publicKey)
	return string(result), err
}

func SaveCertificate(certificate Certifcate) error {
	// сохранение сертификата в базу
	bd, err := mssql.BD()
	if err != nil {
		return err
	}
	defer bd.Close()
	// todo

	return nil
}
