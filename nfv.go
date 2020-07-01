package go_ovh_nfv

import (
	"context"
)

type (
	NFVClient interface {
		Loadbalancer(tenant string) LoadbalancerClient
		Region(tenant string) RegionClient
	}

	RegionClient interface {
		List(context.Context) ([]string, error)
		Get(context.Context, string) (*Region, error)
	}

 	LoadbalancerClient interface {
		List(context.Context) ([]string, error)
		// Required props
		// - region
		// - name
		// - description
		Get(ctx context.Context, id string) (*Loadbalancer, error)
		Create(context.Context, *Loadbalancer) error
		Update(context.Context, *Loadbalancer) error
		Delete(context.Context, string) error

		Configuration(string) ConfigurationClient
	}

	ConfigurationClient interface {
		List(context.Context) ([]int64, error)
		Create(context.Context, *Configuration) error
		Get(context.Context, int64) (*Configuration, error)
		Delete(context.Context, int64) error
		Apply(context.Context, int64) error
	}
)
