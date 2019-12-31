package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "deals-230615")

	if err != nil {
		panic(err)
	}

	Test()

	fs := FlyerService{
		client: client,
	}

	fmt.Print(fs.Flyers())

}
