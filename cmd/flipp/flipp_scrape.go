package main

import (
	"github.com/m-brady/deals/pkg/flipp"
	"log"
)

const SHOPPERS = 208
const LOBLAWS = 2018

func main() {

	response := flipp.GetFlyers()
	flyers := response.Flyers

	for _, flyer := range flyers {
		if flyer.MerchantId == 208 || flyer.MerchantId == 2018 {

		}
	}

	log.Println(flyers)
}
