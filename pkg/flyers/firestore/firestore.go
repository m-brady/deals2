package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/m-brady/deals/pkg/flipp"
	"google.golang.org/api/iterator"
	"log"
	"strconv"
)

func Test() {
	fmt.Println("fsdfs")
}

type FlyerService struct {
	client *firestore.Client
}

func (fs FlyerService) Flyer(id int) (*flipp.Flyer, error) {
	doc, err := fs.client.Collection("flyers").Doc(strconv.Itoa(id)).Get(context.Background())

	if err != nil {
		return nil, err
	}
	data := doc.Data()
	return &flipp.Flyer{
		Id:            data["id"].(int),
		MerchantId:    data["merchant_id"].(int),
		ValidTo:       data["valid_to"].(string),
		ValidFrom:     data["valid_from"].(string),
		AvailableTo:   data["available_to"].(string),
		AvailableFrom: data["available_from"].(string),
	}, nil

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
		data := doc.Data()
		f := &flipp.Flyer{
			Id:            data["id"].(int),
			MerchantId:    data["merchant_id"].(int),
			ValidTo:       data["valid_to"].(string),
			ValidFrom:     data["valid_from"].(string),
			AvailableTo:   data["available_to"].(string),
			AvailableFrom: data["available_from"].(string),
		}
		flyers = append(flyers, f)
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
		data := doc.Data()
		m := &flipp.Merchant{
			Id:             data["id"].(int),
			Name:           data["named"].(string),
			UsBased:        data["us_based"].(bool),
			NameIdentifier: data["name_identifier"].(string),
		}
		merchants = append(merchants, m)
	}
	return merchants, nil
}

func (fs FlyerService) AddFlyerItem(item *flipp.Item) error {
	_, err := fs.client.Collection("item").Doc(strconv.FormatInt(item.Id, 64)).Create(context.Background(), item)
	return err

}
func (fs FlyerService) AddMerchant(merchant *flipp.Merchant) error {
	_, err := fs.client.Collection("item").Doc(strconv.Itoa(merchant.Id)).Create(context.Background(), merchant)
	return err
}
func (fs FlyerService) AddFlyer(flyer *flipp.Flyer) error {
	_, err := fs.client.Collection("item").Doc(strconv.Itoa(flyer.Id)).Create(context.Background(), flyer)
	return err
}
