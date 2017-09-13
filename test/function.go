package main

import (
	"strconv"
)

//translate array to daydata struct
func trsArrTodayData(record []string) (dayData, error) {
	closingprice, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		Error(err)
		closingprice = 0
		return dayData{}, err
	}
	topprice, err := strconv.ParseFloat(record[4], 64)
	if err != nil {
		Error(err)
		topprice = 0
	}
	floorprice, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		Error(err)
		floorprice = 0
	}
	openingPrice, err := strconv.ParseFloat(record[6], 64)
	if err != nil {
		Error(err)
		openingPrice = 0
	}
	beforeclosing, err := strconv.ParseFloat(record[7], 64)
	if err != nil {
		Error(err)
		beforeclosing = 0
	}
	changeamount, err := strconv.ParseFloat(record[8], 64)
	if err != nil {
		Error(err)
		changeamount = 0
	}
	pricelimit, err := strconv.ParseFloat(record[9], 64)
	if err != nil {
		Error(err)
		pricelimit = 0
	}
	turnoverrate, err := strconv.ParseFloat(record[10], 64)
	if err != nil {
		Error(err)
		turnoverrate = 0
	}
	turnovernumber, err := strconv.ParseFloat(record[11], 64)
	if err != nil {
		Error(err)
		turnovernumber = 0
	}
	turnovermoney, err := strconv.ParseFloat(record[12], 64)
	if err != nil {
		Error(err)
		turnovermoney = 0
	}
	aggregatemarketvalue, err := strconv.ParseFloat(record[13], 64)
	if err != nil {
		Error(err)
		aggregatemarketvalue = 0
	}
	circilationmarketvale, err := strconv.ParseFloat(record[14], 64)
	if err != nil {
		Error(err)
		circilationmarketvale = 0
	}
	dealnum, err := strconv.ParseFloat(record[15], 64)
	if err != nil {
		// Error(err)
		dealnum = 0
	}

	d := dayData{
		Date:                   record[0],
		Code:                   record[1],
		Name:                   record[2],
		ClosingPrice:           closingprice,
		TopPrice:               topprice,
		FloorPrice:             floorprice,
		OpeningPrice:           openingPrice,
		BeforeClosing:          beforeclosing,
		ChangeAmount:           changeamount,
		PriceLimit:             pricelimit,
		TurnoverRate:           turnoverrate,
		TurnoverNumber:         turnovernumber,
		TurnoverMoney:          turnovermoney,
		AggregateMarketValue:   aggregatemarketvalue,
		CircilationMarketValue: circilationmarketvale,
		DealNum:                dealnum,
	}
	return d, nil

}
