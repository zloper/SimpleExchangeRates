package main

import "ExchangeRatesRussia/storage"

func main() {
	//TODO refresh every day
	//storage.GetLastMonth()
	//TODO callable
	//storage.GetCurrent("GBP")
	//TODO callable
	storage.GetBestOf("GBP", 30)

	storage.GetBestOf("GBP", 2)

}
