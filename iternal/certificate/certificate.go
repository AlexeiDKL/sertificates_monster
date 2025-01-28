package certificate

import (
	"crypto/rsa"

	"dkl.dklsa.certificates_monster/iternal/cryptoKey"
)

func CreateCertificate(phrase string, publicKey *rsa.PublicKey) (string, error) {
	result, err := cryptoKey.EncryptWithPublicKey(phrase, publicKey)
	return string(result), err
}

func SaveCertificate(certificate string) error {
	// сохранение сертификата в базу

	return nil
}
