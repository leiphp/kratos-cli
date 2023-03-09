package tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDecory(t *testing.T) {
	//saveOneValue("", nil)
	a := math.Ceil(29/10)
	fmt.Println(a)
}

func TestSplite(t *testing.T) {
	////saveOneValue("", nil)
	//str := "2020-07"
	//arr := strings.Split(str, "-")
	//fmt.Println("arr:",arr)
	//aint,_ := strconv.Atoi(arr[1])
	//fmt.Println("val:",aint)
	//获取本月第一天的时间（今天2019-05-20）返回 2019-05-01
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.Format("2006-01-02")
	fmt.Println("monthOneDay:",monthOneDay)

	//获取下个月第一天的时间（今天2019-05-20）返回 2019-06-01
	nextMonthOneDay := thisMonth.AddDate(0, 1, 0).Format("2006-01-02")
	fmt.Println("nextMonthOneDay:",nextMonthOneDay)

}

//随机字符串
func TestRandomString(t *testing.T) {
	n := 6
	//str := "023456789ABCDEFGHJKMNPQRSTUVWXYZ"
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		fmt.Println("i:",i)
		fmt.Println("j:",string(bytes[rand.Intn(len(bytes))]))
		if i == 0 {//第一位不要0
			result = append(result, bytes[rand.Intn(len([]byte("123456789")))])
		}else{
			result = append(result, bytes[rand.Intn(len(bytes))])
		}
		fmt.Println("aaa:",string(result))
	}
	fmt.Println("res:",string(result))
}

func TestMd5(t *testing.T) {
	prolicy := `YCD_{"insureCompanyName":"国任保险","insureCompanyCode":"PCIC","recordNo":"aaaaa","engineeringStartDate":"2022-06-24","engineeringEndDate":"2022-07-24","applicantionNum":1,"applicantName":"爱空间装饰(深圳)有限公司","applicantnCertNo":"91440300358248932N","applicantPhone":"13567555423","projectName":"测试工程","provinceCode":"440000","provinceName":"广东省","cityCode":"440300","cityName":"深圳市","areaCode":"440304","areaName":"福田区","projectAddress":"广东省深圳市南山区粤海街道航城科技大厦2002","projectBuildPrice":"10000"}`
	res := md5Test2(prolicy)
	fmt.Println("res:",res)

	file := "test.txt"
	//s := "hello world"

	Val1, _ := GetFileMd5(file)
	Val2 := GetStringMd5(prolicy)
	fmt.Println("Val1:",Val1)
	fmt.Println("Val2:", Val2)

	h := md5.New()
	h.Write([]byte("123456")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	sig := []byte("123456")
	newSig := md5.Sum(sig)	//转成加密编码
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x",newSig )
	//输出字符串字母都是小写，转换为大写
	sig2 := strings.ToTitle(newArr)
	fmt.Println("sig:",sig2)

	h2 := md5.New()
	h2.Write([]byte("ouba_beatTiger")) // 需要加密的字符串为 ouba_beatTiger
	fmt.Printf("%s\n", hex.EncodeToString(h2.Sum(nil))) // 输出加密结果

}

func md5Test2(str string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5String
}


func GetFileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("os Open error")
		return "", err
	}
	md5 := md5.New()
	_, err = io.Copy(md5, file)
	if err != nil {
		fmt.Println("io copy error")
		return "", err
	}
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str, nil
}

func GetStringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}

func TestRun(t *testing.T){
	//指定分隔符
	str := "32269"
	countSplit := strings.Split(str, ",")
	fmt.Println(countSplit, len(countSplit))
}