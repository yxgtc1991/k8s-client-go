package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

func main() {
	// 1、获取 kubeconfig 路径
	homePath := homedir.HomeDir()
	if homePath == "" {
		log.Fatalf("failed to get the home directory")
	}
	kubeconfig := filepath.Join(homePath, ".kube", "config")
	log.Printf("kube config: %s", kubeconfig)
	// 2、通过 kubeconfig 初始化 config：kubeconfig 已包含 API Server 的连接信息
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	dpClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

	log.Println("create deployment")
	if err = createDeployment(dpClient); err != nil {
		log.Fatal(err)
	}
	<-time.Tick(1 * time.Minute)

	log.Println("update deployment")
	if err = updateDeployment(dpClient); err != nil {
		log.Fatal(err)
	}
	<-time.Tick(1 * time.Minute)

	log.Println("delete deployment")
	if err = deleteDeployment(dpClient); err != nil {
		log.Fatal(err)
	}
	<-time.Tick(1 * time.Minute)
}

func createDeployment(dpClient v1.DeploymentInterface) error {
	replicas := int32(3)
	newDp := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-deploy",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := dpClient.Create(context.TODO(), newDp, metav1.CreateOptions{})
	return err
}

func updateDeployment(dpClient v1.DeploymentInterface) error {
	dp, err := dpClient.Get(context.TODO(), "nginx-deploy", metav1.GetOptions{})
	if err != nil {
		return err
	}
	dp.Spec.Template.Spec.Containers[0].Image = "nginx:1.27.0"
	return retry.RetryOnConflict(
		retry.DefaultRetry, func() error {
			_, err = dpClient.Update(context.TODO(), dp, metav1.UpdateOptions{})
			return err
		},
	)
}

func deleteDeployment(dpClient v1.DeploymentInterface) error {
	deletePolicy := metav1.DeletePropagationForeground
	return dpClient.Delete(context.TODO(), "nginx-deploy", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
