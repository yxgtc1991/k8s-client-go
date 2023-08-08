package deploy_service

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

// UpdateDeployment 更新部署镜像
func UpdateDeployment() {
	deploymentName := flag.String("deployment", "", "deployment name")
	imageName := flag.String("image", "", "new image name")
	appName := flag.String("app", "app", "application name")
	flag.Parse()
	deployment := GetDeployment(deploymentName, imageName)
	name := deployment.GetName()
	fmt.Println("name -> ", name)
	containers := &deployment.Spec.Template.Spec.Containers
	found := false
	// 避坑：通过range方式遍历数组，arr[index]和item的地址值并不是同一个，无法修改原数组！
	/* for _, c := range *containers {
		if c.Name == *appName {
			found = true
			fmt.Println("Old image -> ", c.Image)
			fmt.Println("New image -> ", *imageName)
			c.Image = *imageName
		}
	} */
	c := *containers
	for i := range *containers {
		if c[i].Name == *appName {
			found = true
			fmt.Println("Old image -> ", c[i].Image)
			fmt.Println("New image -> ", *imageName)
			c[i].Image = *imageName
		}
	}
	if found == false {
		fmt.Println("The application container not exist in the deployment pods.")
		os.Exit(0)
	}
	_, err := client_set_service.ClientSet.AppsV1().Deployments("default").Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(deployment.Spec.Template.Spec.Containers[0].Image)
}

// GetDeployment 获取部署信息
func GetDeployment(deploymentName, imageName *string) *v1.Deployment {
	if *deploymentName == "" {
		fmt.Println("you must specify the deployment name.")
		os.Exit(0)
	}
	if *imageName == "" {
		fmt.Println("you must specify the new image name.")
		os.Exit(0)
	}
	clientSet := client_set_service.GetInstance()
	deployment, err := clientSet.AppsV1().Deployments("default").Get(context.TODO(), *deploymentName, metaV1.GetOptions{})
	HandleErr(err)
	fmt.Println("Found deployment")
	return deployment
}

// HandleErr 处理Error
func HandleErr(err error) {
	if errors.IsNotFound(err) {
		fmt.Println("Deployment not found")
	}
	statusErr, isStatus := err.(*errors.StatusError)
	if isStatus {
		fmt.Printf("Error getting deployment:%v\n", statusErr.ErrStatus.Message)
	}
	if err != nil {
		panic(err.Error())
	}
}
