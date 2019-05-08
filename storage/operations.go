package storage

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetLastMonth() {
	ClearStorage()
	for i := 1; i > -30; i-- {
		dt := time.Now().AddDate(0, 0, i)
		curent_date := dt.Format("02/01/2006")
		res, dt_stamp := ReadDailyPrice(curent_date)
		Update(dt_stamp, res)
	}
}

func GetCurrent(cur string) Valute {
	history := GetHistory()
	result := history[LastDate()].Currency[cur]

	fmt.Println(result.Name)
	fmt.Println(result.CharCode)
	fmt.Println(result.Value + " p.")
	return result
}

func GetBestOf(cur string, days int) {

	history := GetHistory()
	maxValue := 0.0
	bestDate := ""
	for i := 1; i > days*(-1); i-- {
		dt := time.Now().AddDate(0, 0, i).Format("02.01.2006")
		strNum := history[dt].Currency[cur].Value
		strNum = strings.Replace(strNum, ",", ".", -1)
		dayValue, _ := strconv.ParseFloat(strNum, 64)
		if maxValue <= dayValue {
			maxValue = dayValue
			bestDate = dt
		}

	}
	fmt.Println("Best price:", maxValue, "("+bestDate+")")

}