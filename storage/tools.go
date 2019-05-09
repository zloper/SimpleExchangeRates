package storage

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch charset {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
}

func ClearStorage() {
	storageFile, err := os.Create("./db.json")
	if err != nil {
		panic(err)
	}
	defer storageFile.Close()

	_, err = storageFile.Write([]byte("{}"))
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON data Updated in file: ", storageFile.Name())
}

func ReadDailyPrice(date string) (map[string]Valute, string) {
	arg := ""
	if date != "" {
		arg = "?date_req=" + date
	}
	resp, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp" + arg)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var okane ValCurs
	// Can't Unmarshal: windows-1251
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = makeCharsetReader
	err = decoder.Decode(&okane)
	if err != nil {
		fmt.Println(err)
	}

	var jishiyo map[string]Valute
	jishiyo = make(map[string]Valute)
	for i := 0; i < len(okane.Valute); i++ {
		jishiyo[okane.Valute[i].CharCode] = okane.Valute[i]
	}

	return jishiyo, okane.Date
}

func GetHistory() Storage {
	history := make(Storage)
	storageFile, err := os.Open("./db.json")
	if err != nil {
		panic(err)
	}
	defer storageFile.Close()

	byteValue, _ := ioutil.ReadAll(storageFile)
	err = json.Unmarshal(byteValue, &history)
	if err != nil {
		panic(err)
	}
	return history
}

func Update(dt string, data map[string]Valute) {
	history := GetHistory()
	var info Info
	info.Currency = data
	history[dt] = info
	//example: fmt.Println(storage["09.05.2019"].Currency["GBP"].Value)

	storageFile, err := os.Create("./db.json")
	if err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(history)
	_, err = storageFile.Write(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON data Updated in file: ", storageFile.Name())
}

func LastDate() string {
	return time.Now().AddDate(0, 0, 1).Format("02.01.2006")
}

func GetLatestDate() string {
	layout := "02.01.2006"

	maxDate, err := time.Parse(layout, "01.01.1900")
	if err != nil {
		panic(err)
	}

	history := GetHistory()
	for key, _ := range history {
		curentDate, err := time.Parse(layout, key)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if maxDate.Unix() < curentDate.Unix() {
			maxDate = curentDate
		}
	}
	return maxDate.Format("02.01.2006")
}
