package client

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestDiscoveryClient(t *testing.T) {
	// 加载 kubeconfig 配置信息
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// 通过kubeconfig配置信息实例化DiscoveryClient对象，该对象用于发现Kubernetes API Server所支持的资源组、资源版本、资源信息的客户端
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	// 获取Kubernetes API Server所支持的资源组、资源版本、资源信息
	_, apiResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	for _, list := range apiResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("group:%v,version:%v\nresources:\n", gv.Group, gv.Version)
		for _, resource := range list.APIResources {
			fmt.Printf("%v\n", resource.Name)
		}
		fmt.Println()
	}
}
