package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	mgo "v1/mgo.v2"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://quotes.money.163.com/service/chddata.html?code=0601398&start=20000720&end=20150508")
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
	fmt.Println(string(bt))
	// strs := make([]byte, 0)

	s := strings.Split(string(bt), "\n")
	Debug(len(s))
	Debug(s[0])
	Debug(s[1])
	ss := strings.Split(s[1], ",")
	Debug(len(ss))
	Debug(ss)
	// strs = append(strs)

}

func main1() {
	err := mongoInit()
	file, err := os.Open("/home/wangjc/下载/601398 (1).csv")
	if err != nil {
		Error(err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)
	var dayDataList []interface{} = make([]interface{}, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		d, err := trsArrTodayData(record)
		if err != nil {
			Error(err)
			continue
		}
		dayDataList = append(dayDataList, d)
	}
	Debug(len(dayDataList))
	Debug(dayDataList[0])
	dataStore(dayDataList)

}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder()) //GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

type mongoConfig struct {
	session *mgo.Session
	db      *mgo.Database
}

var mongo mongoConfig

func mongoInit() error {
	session, err := mgo.Dial("")
	if err != nil {
		return err
	}
	mongo.session = session
	mongo.db = session.DB("Adata")
	return nil
}

//daydata mongo column
func dataStore(arr []interface{}) {
	c := mongo.db.C("daydata")
	Debug(len(arr))
	Debug(arr[0])
	err := c.Insert(arr...)
	if err != nil {
		Error(err)
	}
	Debug("finish")
}
