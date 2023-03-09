package tool

import (
	"fmt"
	"log"
	"testing"
)

func TestGetSm2Keys(t *testing.T) {
	GetSm2Keys()
}

func TestCreateSm2Encrypt(t *testing.T) {
	ciphertxt,pub,err:=CreateSm2Encrypt([]byte("hello world"))
	fmt.Println("ciphertxt:",ciphertxt)
	fmt.Println("ciphertxt:",string(ciphertxt))
	fmt.Println("pub:",pub)
	fmt.Println("err:",err)
	fmt.Println("res:",fmt.Sprintf("%x", ciphertxt))
}

func TestSm2Decrypt(t *testing.T) {
	res, err:=Sm2Decrypt([]byte("3074022100f4ccc01672515b409a3bbd8fbdca089d6f8602142a492dc3c635abf124d28c370220458f3f3a82c75ee8d19410a4c2fa2807ae038ffcdb6c05349c7475266150900b0420c552d1228762f986f688ef3a12e7f5a92af7d97c7460a31b2d7505673755160e040b5a091f6cb7b76bec8da10d"))
	if err != nil {
		log.Fatal("err:",err)
	}
	fmt.Println("res:",res)
}