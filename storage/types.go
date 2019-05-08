package storage

import "encoding/xml"

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []Valute `xml:"Valute"`
	Date    string   `xml:"Date,attr"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  string   `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}

type Info struct {
	Currency map[string]Valute `json:"Currency"`
}

type Storage map[string]Info
