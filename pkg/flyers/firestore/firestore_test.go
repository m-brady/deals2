package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/m-brady/deals/pkg/flipp"
	"testing"
)

func TestAbcd(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "deals-230615")

	if err != nil {
		panic(err)
	}

	fs := FlyerService{
		Client: client,
	}

	_ = fs.AddFlyer(&flipp.Flyer{
		Id:            11111,
		MerchantId:    0,
		ValidTo:       "",
		ValidFrom:     "",
		AvailableTo:   "",
		AvailableFrom: "",
	})

	//fmt.Print(fs.Flyers())

}

func TestAbc(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "deals-230615")

	if err != nil {
		panic(err)
	}

	Test()

	fs := FlyerService{
		client: client,
	}

	f, _ := fs.Flyers()

	fmt.Printf("%v", f)

	//fmt.Print(fs.Flyers())

}
