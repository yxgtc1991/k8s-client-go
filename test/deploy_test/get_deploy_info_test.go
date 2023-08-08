package deploy_test

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/deploy_service"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetDeployInfo(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	deployName := "httpd"
	deployment, err := deploy_service.NewDeployService().GetDeployInfo(ctx, client, namespace, deployName)
	c.Convey("获取部署详情", t, func() {
		c.So(err, c.ShouldEqual, nil)
		c.So(deployment.GetName(), c.ShouldEqual, deployName)
	})
}
