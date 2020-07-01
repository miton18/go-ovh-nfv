package go_ovh_nfv

import (
	"context"
	"fmt"

	"github.com/ovh/go-ovh/ovh"
	"github.com/pkg/errors"
)

func ListRegions(ctx context.Context, client *ovh.Client, tenant string)([]string, error) {
	var regions []string
	url := fmt.Sprintf("/cloud/project/%s/capabilities/loadbalancer/region", tenant)

	if err := client.GetWithContext(ctx, url, &regions); err != nil {
		return nil, err
	}

	return regions, nil
}

func GetRegion(ctx context.Context, client *ovh.Client, tenant, name string) (*Region, error){
	region := &Region{}
	url := fmt.Sprintf("/cloud/project/%s/capabilities/loadbalancer/region/%s", tenant, name)

	if err := client.GetWithContext(ctx, url, region); err != nil {
		return nil, err
	}

	return region, nil
}

func ListLoadbalancer(ctx context.Context, client *ovh.Client, tenant string)([]string, error) {
	var lbs []string
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer", tenant)

	if err := client.GetWithContext(ctx, url, &lbs); err != nil {
		return nil, err
	}

	return lbs, nil
}

func GetLoadbalancer(ctx context.Context, client *ovh.Client, tenant, id string) (*Loadbalancer, error){
	lb := &Loadbalancer{}
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s", tenant, id)

	err := client.GetWithContext(ctx, url, lb)
	if err != nil {
		return nil, err
	}

	return lb, nil
}

func CreateLoadbalancer(ctx context.Context, client *ovh.Client, tenant string, lb *Loadbalancer) error {
	if lb.ID != "" {
		return errors.Errorf("this lb already have an ID, cannot create it")
	}

	url := fmt.Sprintf("/cloud/project/%s/loadbalancer", tenant)
	lbC := &LoadbalancerCreate{
		Region: lb.Region,
		Name: lb.Name,
		Description: lb.Description,
	}

	return client.PostWithContext(ctx, url, lbC, lb)
}

func UpdateLoadbalancer(ctx context.Context, client *ovh.Client, tenant string, loadbalancer *Loadbalancer) error{
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s", tenant, loadbalancer.ID)

	return client.PutWithContext(ctx, url, loadbalancer, loadbalancer)
}

func DeleteLoadbalancer(ctx context.Context, client *ovh.Client, tenant, id string) error{
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s", tenant, id)

	return client.DeleteWithContext(ctx, url, nil)
}

func ListLoadbalancerConfiguration(ctx context.Context, client *ovh.Client, tenant, id string)([]int64, error) {
	var configs []int64
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration", tenant, id)

	if err := client.GetWithContext(ctx, url, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func GetLoadbalancerConfiguration(ctx context.Context, client *ovh.Client, tenant, id string, version int64) (*Configuration, error) {
	config := &Configuration{}
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration/%d", tenant, id, version)

	if err := client.GetWithContext(ctx, url, config); err != nil {
		return nil, err
	}
	return config, nil
}

func CreateLoadbalancerConfiguration(ctx context.Context, client *ovh.Client, tenant, id string, config *Configuration) error{
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration", tenant, id)

	return client.PostWithContext(ctx, url, config, config)
}

func DeleteLoadbalancerConfiguration(ctx context.Context, client *ovh.Client, tenant, id string, version int64) error{
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration/%d", tenant, id, version)

	return client.DeleteWithContext(ctx, url, nil)
}

func ApplyLoadbalancerConfiguration(ctx context.Context, client *ovh.Client, tenant, id string, version int64) error{
	url := fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration/%d/apply", tenant, id, version)

	return client.PostWithContext(ctx, url, nil, nil)
}

