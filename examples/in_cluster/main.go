package main

import (
	"context"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 1、初始化 config：读取认证所需的 token 和 ca.crt
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	// 2、通过 config 初始化 clientset：实现各种资源的 CURD 操作
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	namespace := "default"
	ctx := context.Background()
	for {
		// 3、通过 clientset 列出特定命名空间的所有 Pod
		pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for i, pod := range pods.Items {
			log.Printf("%d -> %s:%s", i+1, pod.Namespace, pod.Name)
		}
		<-time.Tick(5 * time.Second)
	}
}
