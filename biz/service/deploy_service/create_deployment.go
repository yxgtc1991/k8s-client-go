package deploy_service

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateDeploy 创建部署
func (deployServ DeployServiceImpl) CreateDeploy(ctx context.Context, client *kubernetes.Clientset, namespace string, deployment *v1.Deployment) error {
	_, err := client.AppsV1().Deployments(namespace).Create(ctx, deployment, metaV1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
