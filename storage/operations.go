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
	err := rq.ParseForm()
	if err != nil {
		panic(err)
	}
	args := rq.Form

	if rqMsg == "GetCurrent" {
		//help: ...:8087/GetCurrent?cur=GBP
		cur := LstToStr(args["cur"])
		_, rqMsg = GetCurrent(cur)
	} else if rqMsg == "GetLastMonth" {
		GetLastMonth()
		rqMsg = "Get last 30 days exchange rates info"
	} else if rqMsg == "GetBestOf" {
		//help: ...:8087/GetBestOf?cur=GBP&days=14
		cur := LstToStr(args["cur"])
		days, err := strconv.Atoi(LstToStr(args["days"]))
		if err != nil {
			fmt.Println(err)
			rqMsg = fmt.Sprintf("%s", err)
		} else {
			rqMsg = GetBestOf(cur, days)
		}
	} else if rqMsg == "GetAllDays" {
		//help: ...:8087/GetAllDays?cur=GBP
		cur := LstToStr(args["cur"])
		rqMsg = GetAllDays(cur, 7)
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

func GetAllDays(cur string, days int) string {
	history := GetHistory()
	result := ""
	for i := 1; i > days*(-1); i-- {
		dt := time.Now().AddDate(0, 0, i).Format("02.01.2006")
		strNum := history[dt].Currency[cur].Value
		strNum = strings.Replace(strNum, ",", ".", -1)
		strNum = strings.Trim(strNum, " ")
		if strNum != "" {
			fmt.Println(dt)
			result = strNum + "," + result
		}
	}
	result = cur + "=" + result
	result = strings.Trim(result, ",")
	return result
}
