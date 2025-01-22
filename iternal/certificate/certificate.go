package certificate

import "crypto/rsa"

func CreateCertificate(publicKey *rsa.PublicKey) (string, error) {
	// создание сертификата
	// получаем любую не задейственную фразу на сервере
	// формируем сертификат с помощью полученной фразы и публичного ключа
	// возращаем сертификат
	return "", nil // заменить на реальное создание сертификата
}
