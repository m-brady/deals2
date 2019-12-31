package main

import (
	"database/sql"
	"fmt"
	"github.com/m-brady/deals/pkg/flipp"
	"github.com/m-brady/deals/pkg/flyers/mysql"
)

const SHOPPERS = 208
const LOBLAWS = 2018

func main() {

	response := flipp.GetFlyers()
	flyers := response.Flyers
	db, err := sql.Open("mysql", "root:brady@/optimum")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	flyerService := mysql.FlyerService{DB: db}

	for _, flyer := range flyers {
		if flyer.MerchantId == SHOPPERS || flyer.MerchantId == LOBLAWS {
			f, err := flyerService.Flyer(flyer.Id)
			if err != nil {
				flyerService.AddFlyer(&flyer)
			}
			fmt.Println(f)
		}
	}
}
