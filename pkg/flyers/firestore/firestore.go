package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/m-brady/deals/pkg/flipp"
	"google.golang.org/api/iterator"
	"log"
	"strconv"
)

type FlyerService struct {
	client *firestore.Client
}

func (fs FlyerService) Flyer(id int) (*flipp.Flyer, error) {
	doc, err := fs.client.Collection("flyers").Doc(strconv.Itoa(id)).Get(context.Background())

	if err != nil {
		return nil, err
	}
	var flyer flipp.Flyer

	err = doc.DataTo(&flyer)
	if err != nil {
		return nil, err
	}
	return &flyer, nil
}
func (fs FlyerService) Flyers() ([]*flipp.Flyer, error) {

	var flyers []*flipp.Flyer

	iter := fs.client.Collection("flyers").Documents(context.Background())

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		var flyer flipp.Flyer
		err = doc.DataTo(&flyer)
		flyers = append(flyers, &flyer)
	}
	return flyers, nil
}
func (fs FlyerService) Merchants() ([]*flipp.Merchant, error) {
	var merchants []*flipp.Merchant

	iter := fs.client.Collection("merchants").Documents(context.Background())

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		var merchant flipp.Merchant
		err = doc.DataTo(&merchant)
		merchants = append(merchants, &merchant)
	}
	return merchants, nil
}

func (fs FlyerService) AddFlyerItem(item *flipp.Item) error {
	_, err := fs.client.Collection("item").Doc(strconv.FormatInt(item.Id, 64)).Create(context.Background(), item)
	return err

}
func (fs FlyerService) AddMerchant(merchant *flipp.Merchant) error {
	_, err := fs.client.Collection("merchant").Doc(strconv.Itoa(merchant.Id)).Create(context.Background(), merchant)
	return err
}
func (fs FlyerService) AddFlyer(flyer *flipp.Flyer) error {
	_, err := fs.client.Collection("flyer").Doc(strconv.Itoa(flyer.Id)).Create(context.Background(), flyer)
	return err
}
