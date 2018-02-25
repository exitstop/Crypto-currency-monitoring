package lCommon

type ListMonitor struct{
	Index int
	Coin string
	Exchange string
	Price float64
	UpPerPercent float64
	DownPerPercent float64
	UpPer float64
	DownPer float64
	UpLine float64
	DownLine float64
	Hodl float64
	CallBack func(string)(map[string]interface{}, error)
}
