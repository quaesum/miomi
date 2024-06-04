package main

import (
	"log"
	"madmax/services/telegram-scrapper/scrapper"
)

func main() {
	parser := scrapper.New()
	parser.AddSource("https://t.me/s/faunagoroda", 100)
	parser.Run()

	err := parser.Quit()
	if err != nil {
		panic(err)
	}
	log.Println("Parsing successfully finished")
}
