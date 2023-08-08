package deploy

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/deploy_service"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUpdateDeployImage(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	deployName := "httpd"
	image := "2.2.32"
	err := deploy_service.NewDeployService().UpdateDeployImage(ctx, client, namespace, deployName, image)
	c.Convey("更新部署镜像", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}
