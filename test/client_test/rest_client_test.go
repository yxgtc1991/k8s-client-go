package client_test

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestRestClient(t *testing.T) {
	// 加载 kubeconfig 配置信息
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 设置config.APIPath请求的HTTP路径
	config.APIPath = "/api"
	// 设置config.GroupVersion请求的资源组/资源版本
	config.GroupVersion = &coreV1.SchemeGroupVersion
	// 设置config.NegotiatedSerializer数据的编解码器
	config.NegotiatedSerializer = scheme.Codecs

	// 通过kubeconfig配置信息实例化RESTClient对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &coreV1.PodList{}
	// RESTClient对象构建HTTP请求参数
	err = restClient.Get().
		// 设置请求的命名空间
		Namespace("default").
		// 设置请求的资源名称
		Resource("pods").
		// VersionParams函数将一些查询选项（如Limit、TimeoutSeconds等）添加到请求参数中
		VersionedParams(&metaV1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		// 通过Do函数执行请求
		Do(context.TODO()).
		// 将kube-apiserver返回的结果（Result对象）解析到corev1.PodList对象中
		Into(result)
	if err != nil {
		panic(err)
	}
	for _, pod := range result.Items {
		fmt.Printf("namespace:%v,name:%v,status:%v\n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
