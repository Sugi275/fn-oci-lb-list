package ocilib

import (
	"context"

	"github.com/oracle/oci-go-sdk/loadbalancer"
)

// GetLBlist OCIのLB一覧を取得する
func GetLBlist() []string {
	// LB用のクライアントを生成
	client, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(GetEnvConfigProvider())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// LB一覧を取得するためのパラメータを作成
	compartmentid, err := GetCompartmentID()
	if err != nil {
		panic(err)
	}

	request := loadbalancer.ListLoadBalancersRequest{
		CompartmentId: &compartmentid,
	}

	loadbalancers, err := client.ListLoadBalancers(ctx, request)
	if err != nil {
		panic(err)
	}

	var loadbalancersName []string
	for _, loadbalancer := range loadbalancers.Items {
		loadbalancersName = append(loadbalancersName, *loadbalancer.DisplayName)
	}

	return loadbalancersName
}
