package deploy_service

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

type DeployService interface {
	GetDeployInfo(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string) (*v1.Deployment, error)
	UpdateDeployReplica(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string, replicas int32) error
	UpdateDeployReplica2(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string, replicas int32) error
}
