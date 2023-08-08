package deploy

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/deploy_service"
	c "github.com/smartystreets/goconvey/convey"
	v1 "k8s.io/api/apps/v1"
	api "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestCreateDeploy(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()

	namespace := "default"
	replicas := int32(3)
	deployment := &v1.Deployment{}
	deployment.ObjectMeta = metaV1.ObjectMeta{
		Name:      "nginx",
		Namespace: namespace,
		Labels: map[string]string{
			"app": "nginx",
		},
	}
	deployment.Spec = v1.DeploymentSpec{}
	deployment.Spec.Replicas = &replicas
	deployment.Spec.Selector = &metaV1.LabelSelector{
		MatchLabels: map[string]string{
			"app": "nginx",
		},
	}
	deployment.Spec.Template = api.PodTemplateSpec{}
	deployment.Spec.Template.ObjectMeta = metaV1.ObjectMeta{
		Name: "nginx",
		Labels: map[string]string{
			"app": "nginx",
		},
	}
	deployment.Spec.Template.Spec = api.PodSpec{
		Containers: []api.Container{
			{
				Name:  "nginx",
				Image: "nginx:1.16.1",
				Ports: []api.ContainerPort{
					{
						Name:          "http",
						ContainerPort: 80,
						Protocol:      api.ProtocolTCP,
					},
				},
			},
		},
	}
	err := deploy_service.NewDeployService().CreateDeploy(ctx, client, namespace, deployment)
	c.Convey("创建部署", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}
