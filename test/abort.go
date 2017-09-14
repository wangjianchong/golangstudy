package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func csvFile() {
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

//  mongo 创建索引  db.daydata.ensureIndex({"date":-1,"code":-1},{"unique":true})
