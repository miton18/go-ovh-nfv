package go_ovh_nfv

import (
	"context"

	"github.com/ovh/go-ovh/ovh"
)

type (
	nfvClient struct {
		client *ovh.Client
	}

	loadbalancerClient struct {
		client *ovh.Client
		tenant string
	}

	regionClient struct {
		client *ovh.Client
		tenant string
	}

	configurationClient struct {
		client *ovh.Client
		tenant string
		id string
	}
)

func New(client *ovh.Client) NFVClient {
	return &nfvClient{client}
}

func(c *nfvClient) Loadbalancer(tenant string) LoadbalancerClient {
	return &loadbalancerClient{c.client, tenant}
}

func(c *nfvClient) Region(tenant string) RegionClient {
	return &regionClient{c.client, tenant}
}

func (r *regionClient) List(ctx context.Context) ([]string, error) {
	return ListRegions(ctx, r.client, r.tenant)
}

func (r *regionClient) Get(ctx context.Context, region string) (*Region, error) {
	return GetRegion(ctx, r.client, r.tenant, region)
}

func (l *loadbalancerClient) Configuration(id string) ConfigurationClient {
	return &configurationClient{l.client, l.tenant, id}
}

func (l *loadbalancerClient) List(ctx context.Context) ([]string, error) {
	return ListLoadbalancer(ctx, l.client, l.tenant)
}

func (l *loadbalancerClient) Get(ctx context.Context, id string) (*Loadbalancer, error) {
	return GetLoadbalancer(ctx, l.client, l.tenant, id)
}

func (l *loadbalancerClient) Create(ctx context.Context, lb *Loadbalancer) error {
	return CreateLoadbalancer(ctx, l.client, l.tenant, lb)
}

func (l *loadbalancerClient) Update(ctx context.Context, lb *Loadbalancer) error {
	return UpdateLoadbalancer(ctx, l.client, l.tenant, lb)
}

func (l *loadbalancerClient) Delete(ctx context.Context, id string) error {
	return DeleteLoadbalancer(ctx, l.client, l.tenant,  id)
}

func (c *configurationClient) List(ctx context.Context) ([]int64, error) {
	return ListLoadbalancerConfiguration(ctx, c.client, c.tenant, c.id)
}

func (c *configurationClient) Create(ctx context.Context, config *Configuration) error {
	return CreateLoadbalancerConfiguration(ctx, c.client, c.tenant, c.id, config)
}

func (c *configurationClient) Get(ctx context.Context, version int64) (*Configuration, error) {
	return GetLoadbalancerConfiguration(ctx, c.client, c.tenant, c.id, version)
}

func (c *configurationClient) Delete(ctx context.Context, version int64) error {
	return DeleteLoadbalancerConfiguration(ctx, c.client, c.tenant, c.id, version)
}

func (c *configurationClient) Apply(ctx context.Context, version int64) error {
	return ApplyLoadbalancerConfiguration(ctx, c.client, c.tenant, c.id, version)
}
