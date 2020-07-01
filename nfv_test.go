package go_ovh_nfv

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ovh/go-ovh/ovh"
)

var (
	client *ovh.Client
	tenant string
)

func init() {
	tenant = os.Getenv("TENANT")

	var err error
	client, err = ovh.NewClient(
		ovh.OvhEU,
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		os.Getenv("CONSUMER_KEY"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Test(t *testing.T) {
	nfv := New(client)
	ctx := context.Background()
	region := ""
	lb := &Loadbalancer{
		Description: "test-remi",
		Name: "test-remi",
	}

	t.Run("List regions", func(tt *testing.T) {
		regions, err := nfv.Region(tenant).List(ctx)
		if err != nil {
			t.Error("cannot list regions", err)
		}
		fmt.Println(fmt.Sprintf("%+v", regions))

		region = regions[0]
	})

	t.Run("Get region", func(tt *testing.T) {
		region, err := nfv.Region(tenant).Get(ctx, region)
		if err != nil {
			t.Error("cannot list regions", err)
		}
		fmt.Println(fmt.Sprintf("%+v", region))
	})

	t.Run("Create Loadbalancer", func(tt *testing.T) {
		lb.Region = region

		err := nfv.Loadbalancer(tenant).Create(ctx, lb)
		if err != nil {
			t.Error("cannot create lb", err)
		}
		fmt.Println(fmt.Sprintf("%+v", lb))

		if lb.ID == "" {
			t.Error("expect an ID as result of createLB()")
		}
	})

	t.Run("List Loadbalancer", func(tt *testing.T) {
		lbs, err := nfv.Loadbalancer(tenant).List(ctx)
		if err != nil {
			t.Error("cannot list lbs", err)
		}
		fmt.Println(fmt.Sprintf("%+v", lbs))

		if len(lbs) == 0 {
			t.Error("expect at least 1 LB")
		}
	})

	t.Run("Get Loadbalancer", func(tt *testing.T) {
		_lb, err := nfv.Loadbalancer(tenant).Get(ctx, lb.ID)
		if err != nil {
			t.Error("cannot list lbs", err)
		}
		fmt.Println(fmt.Sprintf("%+v", _lb))
	})

	t.Run("Create config", func(tt *testing.T) {
		//TODO: test configs calls
	})

	t.Run("Delete Loadbalancer", func(tt *testing.T) {
		err := nfv.Loadbalancer(tenant).Delete(ctx, lb.ID)
		if err != nil {
			t.Error("cannot delete lb", err)
		}
	})
}

