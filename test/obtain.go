package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	// todayGetDayData("1300083")
	firstGetDayData("1300083")
}

//第一次获取此股信息 调用接口
func firstGetDayData(code string) {
	getDayData(fmt.Sprintf("http://quotes.money.163.com/service/chddata.html?code=%s&start=20000720&end=20180508", code))
}

func obDataByList() {
	stockList := getFreeStock()
	for _, v := range stockList {
		todayGetDayData(v)
	}
}

//定时功能 定时调用
func timeTicker() {
	ticker := time.NewTicker(time.Minute * 10)
	for range ticker.C {
		if time.Now().Format("15") == "15" {
			obDataByList()
		}
	}
}

//每天在三点准时调用这个接口获取当天数据
func todayGetDayData(code string) {
	today := time.Now().Format("20060102")
	Debug(today)
	today = "20170913"
	getDayData(fmt.Sprintf("http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=%s", code, today, today))

}

//将得到的信息保存到mongo中
func getDayData(url string) {
	//0开头应该沪市股票 1开头是深市 创业板股票  fmt.Sprintf("http://quotes.money.163.com/service/chddata.html?code=%s&start=20000720&end=20180508", code)
	resp, err := http.Get(url)
	if err != nil {
		Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	bt, err := GbkToUtf8(body)
	if err != nil {
		Error(err)
		return
	}
	s := strings.Split(string(bt), "\n")
	Debug(len(s))
	arr := make([]interface{}, 0)
	for _, v := range s {
		ss := strings.Split(v, ",")
		if len(ss) != 16 {
			Debug(ss)
			continue
		}
		d, err := trsArrTodayData(ss)
		if err != nil {
			Error(err)
			continue
		}
		arr = append(arr, d)
	}
	dataStore(arr)
}
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder()) //GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//daydata mongo column
func dataStore(arr []interface{}) {
	c := mongo.db.C("daydata")
	Debug(len(arr))
	if len(arr) == 0 {
		Debug("mongo store num :", len(arr))
		return
	}
	Debug(arr[0])
	err := c.Insert(arr...)
	if err != nil {
		Error(err)
	}
	Debug("finish")
}
