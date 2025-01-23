package certificate

import "crypto/rsa"

func CreateCertificate(phrase string, publicKey *rsa.PublicKey) (string, error) {
	// создание сертификата
	// формируем сертификат с помощью полученной фразы и публичного ключа
	// возращаем сертификат
	return "", nil // заменить на реальное создание сертификата
}

func SaveCertificate(certificate string) error {
	// сохранение сертификата в базу

	return nil
}
