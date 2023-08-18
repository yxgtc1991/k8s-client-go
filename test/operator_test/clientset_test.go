package operator_test

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestClientset(t *testing.T) {

	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// client
	client := clientset.CoreV1()

	// get data
	pod, err := client.Pods("default").Get(context.TODO(), "mypod", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	println(pod.Name)
}
