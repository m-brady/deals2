package mysql

import (
	"database/sql"
	"github.com/m-brady/deals/pkg/flipp"
	"log"
)

type FlyerService struct {
	DB *sql.DB
}

func (fs FlyerService) Flyer(id int) (*flipp.Flyer, error) {
	var flyer flipp.Flyer
	err := fs.DB.QueryRow("select id, merchant_id, valid_to, valid_from, available_to, available_from from optimum.flyer ").Scan(flyer)
	if err != nil {
		return nil, err
	} else {
		return &flyer, nil
	}
}
func (fs FlyerService) Flyers() ([]*flipp.Flyer, error) {

	var flyers []*flipp.Flyer

	rows, err := fs.DB.Query("select id, merchant_id, valid_to, valid_from, available_to, available_from from optimum.flyer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var flyer flipp.Flyer
		if err := rows.Scan(&flyer); err != nil {
			log.Fatal(err)
		}
		flyers = append(flyers, &flyer)
	}
	return flyers, nil
}
func (fs FlyerService) Merchants() ([]*flipp.Merchant, error) {
	var merchants []*flipp.Merchant

	rows, err := fs.DB.Query("select id, us_based, name_identifier, name from optimum.merchant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var merchant flipp.Merchant
		if err := rows.Scan(&merchant); err != nil {
			log.Fatal(err)
		}
		merchants = append(merchants, &merchant)
	}
	return merchants, nil
}

func (fs FlyerService) AddFlyerItem(item *flipp.Item) error {

	_, err := fs.DB.Exec("insert ignore into flyer_item (flyer_item_id, name, flyer_id, current_price, valid_to, valid_from) values (?, ?, ?, ?, ? ,?)",
		item.Id, item.Name, item.FlyerId, item.Price, item.ValidTo, item.ValidFrom)

	return err

}
func (fs FlyerService) AddMerchant(merchant *flipp.Merchant) error {
	_, err := fs.DB.Exec("insert ignore into merchant (id, us_based, name_identifier, name) values (?, ?, ?, ?)",
		merchant.Id, merchant.UsBased, merchant.NameIdentifier, merchant.Name)
	return err
}
func (fs FlyerService) AddFlyer(flyer *flipp.Flyer) error {

	_, err := fs.DB.Exec("insert ignore into flyer (id, merchant_id, valid_from, valid_to, available_from, available_to) values (?, ?, ?, ?, ? ,?)",
		flyer.Id, flyer.MerchantId, flyer.ValidFrom, flyer.ValidTo, flyer.AvailableFrom, flyer.AvailableTo)

	return err

}
