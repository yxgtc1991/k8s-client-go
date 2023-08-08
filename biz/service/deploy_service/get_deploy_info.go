package deploy_service

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetDeployInfo 获取部署详情
func (deployServ DeployServiceImpl) GetDeployInfo(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string) (deployment *v1.Deployment, err error) {
	deployment, err = client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	fmt.Println("deployment name => ", deployment.GetName())
	return deployment, nil
}
