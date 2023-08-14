package client

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestDynamicClient(t *testing.T) {
	// 加载 kubeconfig 配置信息
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 通过kubeconfig配置信息实例化DynamicClient对象，该对象用于管理Kubernetes的所有Resource的客户端
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unstructObj, err := dynamicClient.
		// 设置请求的资源组、资源版本、资源名称
		Resource(gvr).
		// 设置请求的命名空间
		Namespace(coreV1.NamespaceDefault).
		// 获取Pod列表，得到的Pod列表为Unstructured.UnstructuredList指针类型
		List(context.TODO(), metaV1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	podList := &coreV1.PodList{}
	// 将Unstructured.UnstructuredList转换成PodList类型
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(
		unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}
	for _, pod := range podList.Items {
		fmt.Printf("namespace:%v,name:%v,status:%v\n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
