package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	once sync.Once
	Token string
)

const (
	host = "http://127.0.0.1:9009"
	clientId = "7191366a52684669a67fc2073d0d5f72"
	secretKey = "14fd036422d446ea874ce98616810076"
	grantType = "authorization_code"
)

//核酸数据
type NucleicData struct {
	Idno          string  `json:"idno"`
	Name          string  `json:"name"`
	Cysj          string  `json:"cysj"`
	Jcsj          string  `json:"jcsj"`
	Jcjg          string  `json:"jcjg"`
	Jcjieguo      string  `json:"jcjieguo"`
}



func CreatFile(){
	channel := make(chan int,1)
	obj,_ := NewStoreDB()
	obj.initStore()
	time.Sleep(1000*time.Second)
	obj.closeCh<- struct{}{} //写个空struct,控制停止定时任务
	<-channel
}

type storeDB struct {
	paramConfig       map[string]interface{}
	productsBarcodes  map[string]string //商品二维码
	kdsProgress       bool
	//外卖相关处理的锁
	takeoutLock sync.Mutex
	ticker      *time.Ticker
	closeCh     chan struct{}
}

func NewStoreDB() (*storeDB, error) {
	s := &storeDB{
		paramConfig:  make(map[string]interface{}),
		closeCh:      make(chan struct{}),
		//ticker:       time.NewTicker(1 * time.Second), //可以不定义ticker,只要返回*storeDB就代表初始化了，后面使用d.ticker和d.closeCh才不会经典报错
	}
	return s, nil
}

//拉取完主档之后的门店初始化操作
//包括初始化gateway连接，发布硬件信息等等
func (d *storeDB) initStore() {
	//fmt.Println("::::",d.ticker)
	if d.ticker == nil {
		d.ticker = time.NewTicker(5 * time.Second)
		//ticker := time.NewTicker(time.Second * 1)
		go func() {
			for {
				select {
				case <-d.ticker.C:
					file, err := os.Open("hesun.txt")
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()
					fl,_ := os.OpenFile("hesun.log", os.O_APPEND, 0777) //已追加的方式打开filepath路径下的文件
					defer fl.Close()
					scanner := bufio.NewScanner(file)
					for scanner.Scan() {
						lineText := scanner.Text()
						fmt.Println("lineText",lineText)
						res,err := d.syncTakeout(lineText)
						if err != nil {
							log.Printf("同步核酸失败", err)
						}
						res.Idno = lineText
						b, err := json.Marshal(res)
						fmt.Println("str:",string(b))
						fl.Write([]byte(b))             //data是自己定义的数据
						fl.WriteString("\n")
						time.Sleep(5*time.Second)
					}

				case <-d.closeCh:
					log.Printf("定时任务退出")
					return
				}
			}
		}()
	}

	//d.addCronTask("外卖同步", "0 */2 * * * *", d.syncTakeout)
	//d.scheduler.Start()
	//d.addCronTask("定时补拉外卖", "0 */5 * * * *", d.syncTakeout)
}

func (d *storeDB) syncTakeout(idno string) (NucleicData,error) {
	fmt.Println("正在执行同步核酸任务")
	res,err :=GetNucleicTestTime(idno)
	if err != nil {
		return NucleicData{idno,"无记录","", "", "", ""},err
	}
	fmt.Println("res:",res)
	res = NucleicData{idno,"无记录","111222", "222", "", ""}
	return res,nil
}

func GetNucleicTestTime(idcard_no string) (NucleicData, error){
	var nucleicData NucleicData
	type bodyJson struct {
		Msg          string  `json:"msg"`                //消息
		Code         int     `json:"code"`               //状态码
		Data         []NucleicData   `json:"data"`       //数据
	}

	client := &http.Client{}
	//生成要访问的url
	url :=  host+"/proxy/rest/hesuan/list"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header["token"] = []string{"dd3b0741dbd58e94603028235ef344b98d31460e"}
	request.Header["clientId"] = []string{clientId}
	request.Header["secretKey"] = []string{secretKey}

	q := request.URL.Query()
	q.Add("idno", idcard_no)
	request.URL.RawQuery = q.Encode()

	resp, err := client.Do(request)
	if err != nil {
		return nucleicData, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nucleicData, err
	}
	var info bodyJson
	json.Unmarshal(bs, &info)
	if info.Code == 200 {
		if len(info.Data) > 0 {
			nucleicData = info.Data[0]
		}else{
			nucleicData.Name = "无记录"
			return nucleicData, errors.New("核酸数据不存在")
		}
	}
	return nucleicData, nil
}