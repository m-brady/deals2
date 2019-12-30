package flyers

import (
	"github.com/m-brady/deals/pkg/flipp"
)

type FlyerService interface {
	Flyer(id int) (*flipp.Flyer, error)
	Flyers() ([]*flipp.Flyer, error)
	Merchants() ([]*flipp.Merchant, error)

	AddFlyerItem(item *flipp.Item) error
	AddMerchant(merchant *flipp.Merchant) error
	AddFlyer(flyer *flipp.Flyer) error
}
