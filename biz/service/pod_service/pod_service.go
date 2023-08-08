package pod_service

import (
	"context"
	"k8s.io/client-go/kubernetes"
)

type PodService interface {
	GetPodInfo(ctx context.Context, client *kubernetes.Clientset, namespace string) error
}
