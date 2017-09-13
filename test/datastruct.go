package main

//dayData struct
type dayData struct {
	Date                   string  "bson:`date`"                   //日期
	Code                   string  "bson:`code`"                   //股票代码
	Name                   string  "bson:`name`"                   //名称
	ClosingPrice           float64 "bson:`closingprice`"           //收盘价
	TopPrice               float64 "bson:`topprice`"               //最高价
	FloorPrice             float64 "bson:`floorprice`"             //最低价
	OpeningPrice           float64 "bson:`openingprice`"           //开盘价
	BeforeClosing          float64 "bson:`beforeclosing`"          //前收盘
	ChangeAmount           float64 "bson:`changeamount`"           //涨跌额
	PriceLimit             float64 "bson:`pricelimit`"             //涨跌幅
	TurnoverRate           float64 "bson:`turnoverrate`"           //换手率
	TurnoverNumber         float64 "bson:`turnovernumber`"         //成交量
	TurnoverMoney          float64 "bson:`turnovermoney`"          //成交金额
	AggregateMarketValue   float64 "bson:`aggregatemarketvalue`"   //总市值
	CircilationMarketValue float64 "bson:`circilationmarketvalue`" //流通市值
	DealNum                float64 "bson:`dealnum`"                //成交笔数
}
