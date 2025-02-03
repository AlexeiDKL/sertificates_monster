package main

import (
	"fmt"
	"reflect"

	"dkl.dklsa.certificates_monster/iternal/certificate"
	"dkl.dklsa.certificates_monster/iternal/config"
	"dkl.dklsa.certificates_monster/iternal/cryptoKey"
	"dkl.dklsa.certificates_monster/iternal/logger"
	storage "dkl.dklsa.certificates_monster/iternal/storage/mssql"
)

func initializing() {
	config.Init()
	logger.Init()
}

func main() {
	initializing()

	phrase, err := storage.GetPhrase()
	if err != nil {
		fmt.Println("phrase error: ", err)
		return
	}

	PKeys, err := cryptoKey.GetPrivateKey("key.secret")
	if err != nil {
		fmt.Println("private key error: ", err)
		return
	}
	certificate, err := certificate.CreateCertificate(phrase, &PKeys.PublicKey)
	if err != nil {
		fmt.Println("certificate error: ", err)
		return
	}

	rr, err := cryptoKey.DecryptWithPrivateKey(certificate, PKeys)

	if err != nil {
		fmt.Println("decrypt error: ", err)
		return
	}
	fmt.Println(reflect.DeepEqual(rr, phrase))

	storage.WG.Wait()
}
