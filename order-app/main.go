package main

import (
	"log"

	"github.com/ysfkel/order-app/boot"
)

func main() {

	err := boot.Start()

	if err != nil {
		log.Fatalln(err)
	}

}
