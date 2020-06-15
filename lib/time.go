package lib

import "time"

//获取当月1号文本时间
func MonthOneDay() string {
	year, month, _ := time.Now().Date()
	monthOneDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	return monthOneDay
}

//获取当月最后一天文本时间
func MonthLastDay() string {
	year, month, _ := time.Now().Date()
	MonthLastDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Format("2006-01-02 15:04:05")
	return MonthLastDay
}

//获取上个月1号文本时间
func LastMonthOneDay() string {
	year, month, _ := time.Now().Date()
	lastMonthOneDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
	return lastMonthOneDay
}

//获取当月1号时间戳
func MonthOneDayUnix() int64 {
	year, month, _ := time.Now().Date()
	monthOneDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
	return monthOneDay
}

//获取上个月1号时间戳
func LastMonthOneDayUnix() int64 {
	year, month, _ := time.Now().Date()
	lastMonthOneDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, -1, 0).Unix()
	return lastMonthOneDay
}

//获取当月最后一天时间戳
func MonthLastDayUnix() int64 {
	year, month, _ := time.Now().Date()
	MonthLastDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Unix()
	return MonthLastDay
}

//获取当前时间
func CurrentTime() string {
	CurrentTime := time.Now().Format("2006-01-02 15:04:05")
	return CurrentTime
}
