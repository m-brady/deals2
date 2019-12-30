package flyers

import (
	"github.com/m-brady/deals/pkg/flipp"
)

type FlyerService interface {
	Flyer(id int) (*Flyer, error)
	CreateFlyer(f *Flyer) error
	Flyers() ([]*Flyer, error)

	Merchants() ([]*Merchant, error)
}
