package deploy_service

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// UpdateDeployReplica 更新部署副本数
func (deployServ DeployServiceImpl) UpdateDeployReplica(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string, replicas int32) error {
	deployment, err := deployServ.GetDeployInfo(ctx, client, namespace, deployName)
	if err != nil {
		return err
	}
	deployment.Spec.Replicas = &replicas
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	fmt.Println("deployment replicas => ", *deployment.Spec.Replicas)
	return nil
}

// UpdateDeployReplica2 更新部署副本数
func (deployServ DeployServiceImpl) UpdateDeployReplica2(ctx context.Context, client *kubernetes.Clientset, namespace, deployName string, replicas int32) error {
	replica, err := client.AppsV1().Deployments(namespace).GetScale(ctx, deployName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	replica.Spec.Replicas = replicas
	replica, err = client.AppsV1().Deployments(namespace).UpdateScale(ctx, deployName, replica, metav1.UpdateOptions{})
	fmt.Println("replica name => ", replica.Name)
	return nil
}
