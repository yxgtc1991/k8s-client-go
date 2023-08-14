package client

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestClientSet(t *testing.T) {
	// 加载 kubeconfig 配置信息
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 通过kubeconfig配置信息实例化ClientSet对象，该对象用于管理所有Resource的客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// clientset.CoreV1().Pods表示请求core资源组的v1资源版本下的Pod资源对象
	// Pods函数是一个资源接口对象，用于Pod资源对象的管理
	podClient := clientset.CoreV1().Pods(coreV1.NamespaceDefault)

	// podClient.List函数通过RESTClient获得Pod列表
	result, err := podClient.List(context.TODO(), metaV1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	for _, pod := range result.Items {
		fmt.Printf("namespace:%v,name:%v,status:%v\n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
