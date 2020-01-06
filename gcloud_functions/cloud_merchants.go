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

func Merchants(ctx context.Context, m PubSubMessage) error {

	merchants := flipp.GetMerchants().Merchants

	client, err := firestore.NewClient(context.Background(), "deals-230615")

	if err != nil {
		panic(err)
	}

	fs := firestoreFlyers.FlyerService{Client: client}
	for _, merchant := range merchants {
		fs.AddMerchant(&merchant)

	}

	return nil
}
