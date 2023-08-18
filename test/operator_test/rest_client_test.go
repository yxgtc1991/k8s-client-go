package operator_test

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestRestClient(t *testing.T) {

	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 必填项
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	pod := v1.Pod{}
	err = restClient.Get().
		Namespace("default").
		Resource("pods").
		Name("mypod").
		Do(context.TODO()).
		Into(&pod)
	if err != nil {
		print(err)
		return
	}
	println(pod.Name)
}
