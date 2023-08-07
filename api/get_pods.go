package api

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods 获取Pod列表
func GetPods() {
	clientSet := GetClientSet()
	pods, err := clientSet.CoreV1().Pods("default").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(pods.Items))
}
