package pod_service

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetPodInfo 获取Pod详情
func (podServ PodServiceImpl) GetPodInfo(ctx context.Context, client *kubernetes.Clientset, namespace string) error {
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Println("pod name => ", pods.Items[0].Status.ContainerStatuses[0].Name)
	fmt.Println("pod image => ", pods.Items[0].Status.ContainerStatuses[0].Image)
	fmt.Println("pod state => ", pods.Items[0].Status.ContainerStatuses[0].State.Running)
	return nil
}
