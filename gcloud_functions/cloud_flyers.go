package gcloudFunctions

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/m-brady/deals/pkg/flipp"
	firestoreFlyers "github.com/m-brady/deals/pkg/flyers/firestore"
)

// HelloPubSub consumes a Pub/Sub message.
//func Merchants(ctx context.Context, m PubSubMessage) error {
//
//	response := flipp.GetFlyers()
//	flyers := response.Flyers
//	client, err := firestore.NewClient(context.Background(), "deals-230615")
//
//	if err != nil {
//		panic(err)
//	}
//
//	fs := firestoreFlyers.FlyerService{Client: client}
//
//	for _, flyer := range flyers {
//		if flyer.MerchantId == SHOPPERS || flyer.MerchantId == LOBLAWS {
//			f, err := fs.Flyer(flyer.Id)
//			if err != nil {
//				fs.AddFlyer(&flyer)
//			}
//			fmt.Println(f)
//		}
//	}
//	return nil
//}

func Flyers(ctx context.Context, m PubSubMessage) error {

	validMerchants := map[int]bool{LOBLAWS: true, SHOPPERS: true, SOBEYS: true, METRO: true}

	flyers := flipp.GetFlyers().Flyers

	client, err := firestore.NewClient(context.Background(), "deals-230615")

	if err != nil {
		panic(err)
	}

	fs := firestoreFlyers.FlyerService{Client: client}
	for _, flyer := range flyers {
		if _, ok := validMerchants[flyer.MerchantId]; ok {

			flipp.GetFlyer(flyer.Id)

			fs.AddFlyer(&flyer)
		}
	}

	return nil
}
