package util

import (
	"math/rand"
	"time"
)

// RandomString 生成n位随机字符串
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().UnixMicro())
	for i := range result {
		result[i] += letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// StringMonthToTime 字符串格式的时间转换为时间对象
func StringMonthToTime(toBeCharge string) (time.Time, error) {
	timeLayout := "2006/1"                 //转化所需模板
	loc, err := time.LoadLocation("Local") //重要：获取时区

	if err != nil {
		return time.Time{}, err
	}

	theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return time.Time{}, err
	}

	return theTime, nil
}
