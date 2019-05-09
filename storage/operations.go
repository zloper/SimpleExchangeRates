package storage

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FuncProvider(writer http.ResponseWriter, rq *http.Request) {
	rqMsg := strings.TrimPrefix(rq.URL.Path, "/")

	if rqMsg == "GetCurrent" {
		_, rqMsg = GetCurrent("GBP")
	} else if rqMsg == "GetLastMonth" {
		GetLastMonth()
	} else if rqMsg == "GetBestOf" {
		rqMsg = GetBestOf("GBP", 14)
	}

	if _, err := writer.Write([]byte(rqMsg)); err != nil {
		panic(err)
	}
}

func GetLastMonth() {
	ClearStorage()
	for i := 1; i > -30; i-- {
		dt := time.Now().AddDate(0, 0, i)
		curent_date := dt.Format("02/01/2006")
		res, dt_stamp := ReadDailyPrice(curent_date)
		Update(dt_stamp, res)
	}
}

func GetCurrent(cur string) (Valute, string) {
	history := GetHistory()
	result := history[GetLatestDate()].Currency[cur]

	strResult := "[" + GetLatestDate() + "]\n"
	strResult += result.Name + "\n"
	strResult += result.CharCode + "\n"
	strResult += result.Value + " p." + "\n"
	fmt.Println(strResult)
	return result, strResult
}

func GetBestOf(cur string, days int) string {
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
	strFloat := strconv.FormatFloat(maxValue, 'f', -1, 64)
	result := "Best price:" + strFloat + "(" + bestDate + ")"
	return result

}
