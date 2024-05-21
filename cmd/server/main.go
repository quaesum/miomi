package main

import (
	"log"
	"madmax/internal"
)

func main() {

	if err := internal.Run(); err != nil {
		log.Fatal(err)
	}

}
