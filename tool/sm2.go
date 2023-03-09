package tool

import (
"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
"github.com/tjfoc/gmsm/sm2"
"log"
"math/big"
)

/*
	用于生成sm2密钥对
	密钥对将会存储在sm2Value
*/
func GetSm2Keys()  {
	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		log.Fatal(err)
	}
	pub := &priv.PublicKey

	var values []*big.Int
	values = append(values, priv.D,pub.X,pub.Y)
	saveValue("sm2Value.txt",values)
}

/*
	使用对方的公钥加密
	content是需要加密的内容
*/
func CreateSm2Encrypt(msg []byte) ([]byte,*sm2.PublicKey,error) {
	//读取密钥对

	re := ReadValue("sm2Value.txt")

	if len(re)<3 {
		return nil,nil,errors.New("文件内容损坏！")
	}

	c := sm2.P256Sm2()
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = new(big.Int).Set(re[0])
	priv.PublicKey.X, priv.PublicKey.Y = new(big.Int).Set(re[1]),new(big.Int).Set(re[2])

	pub := &priv.PublicKey
	ciphertxt, err := pub.EncryptAsn1(msg,rand.Reader) //sm2加密
	if err != nil {
		return nil, nil, err
	}

	return ciphertxt,pub,nil
}

/*
	私钥进行解密操作
	需要使用匹配的私钥进行解密
*/
func Sm2Decrypt(ciphertxt []byte) (string,error) {
	//读取密钥对
	re := ReadValue("sm2Value.txt")
	fmt.Println("re:",re)

	if len(re)<3 {
		return "", errors.New("文件内容损坏！")
	}

	c := sm2.P256Sm2()
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = new(big.Int).Set(re[0])
	priv.PublicKey.X, priv.PublicKey.Y = new(big.Int).Set(re[1]),new(big.Int).Set(re[2])

	plaintxt,err :=  priv.DecryptAsn1(ciphertxt)  //sm2解密
	if err != nil {
		return "", err
	}

	return string(plaintxt),nil
}

/*
	使用私钥创建签名
*/
func CreateSm2Sig(msg []byte) ([]byte,*sm2.PublicKey,error) {
	//读取密钥对
	re := ReadValue("sm2Value.txt")

	if len(re)<3 {
		return nil,nil, errors.New("文件内容损坏！")
	}

	c := sm2.P256Sm2()
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = new(big.Int).Set(re[0])
	priv.PublicKey.X, priv.PublicKey.Y = new(big.Int).Set(re[1]),new(big.Int).Set(re[2])

	sign,err := priv.Sign(rand.Reader, msg, nil)  //sm2签名
	if err != nil {
		return nil,nil, err
	}
	return sign,&priv.PublicKey,err
}

/*
	验证签名是否正常
*/
func VerSm2Sig(pub *sm2.PublicKey,msg []byte,sign []byte) bool {
	isok := pub.Verify(msg, sign)
	if !isok {
		return false
	}
	return true
}



