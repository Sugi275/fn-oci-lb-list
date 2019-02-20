package ocilib

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/loadbalancer"
)

// GetLBlist OCIのLB一覧を取得する
func GetLBlist() []string {
	// LB用のクライアントを生成
	//client, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(GetEnvConfigProvider()
	client, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	debugprint(GetEnvConfigProvider())
	debugprint(common.DefaultConfigProvider())

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

func debugprint(c common.ConfigurationProvider) {
	s1, _ := c.PrivateRSAKey()
	fmt.Println("PrivateRSAKey", s1)

	s2, _ := c.KeyID()
	fmt.Println("KeyID", s2)

	s3, _ := c.TenancyOCID()
	fmt.Println("TenancyOCID", s3)

	s4, _ := c.UserOCID()
	fmt.Println("UserOCID", s4)

	s5, _ := c.KeyFingerprint()
	fmt.Println("KeyFingerprint", s5)

	s6, _ := c.Region()
	fmt.Println("Region", s6)

}
