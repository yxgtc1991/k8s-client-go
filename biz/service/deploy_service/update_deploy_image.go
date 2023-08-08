package deploy_service

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// UpdateDeployImage 更新部署镜像
func (deployServ DeployServiceImpl) UpdateDeployImage(ctx context.Context, client *kubernetes.Clientset, namespace, deployName, image string) error {
	deployment, err := deployServ.GetDeployInfo(ctx, client, namespace, deployName)
	if err != nil {
		return err
	}
	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metaV1.UpdateOptions{})
	fmt.Println("image name => ", deployment.Spec.Template.Spec.Containers[0].Image)
	return nil
}
