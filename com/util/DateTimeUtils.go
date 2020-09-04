package util

import "time"

// 获取当前时间
func GetNowUnix() int64 {
	return time.Now().Unix()
}

// 获取月初时间
func GetEarlyMonthUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取零时时间
func GetZeroHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取当前小时时间
func GetNowHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取年初时间
func GetEarlyYearUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

/**
格式化日期
*/
func GetUnixToFormatString(timestamp int64, f string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(f)
}

/**
当前日期 2006-01-02
*/
func GetToday() string {
	today := GetUnixToFormatString(GetNowUnix(), "2006-01-02")
	return today
}

/**
获取当前日期
*/
func GetTodayTime() string {
	todayTime := GetUnixToFormatString(GetNowUnix(), "2006-01-02 15:04:05")
	return todayTime
}
