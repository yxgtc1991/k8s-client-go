package deploy_service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeleteDeploy 删除部署
func (deployServ DeployServiceImpl) DeleteDeploy(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string) error {
	err := client.AppsV1().Deployments(namespace).Delete(ctx, deployName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
