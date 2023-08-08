package deploy

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/deploy_service"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDeleteDeploy(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	deployName := "my-nginx"
	err := deploy_service.NewDeployService().DeleteDeploy(ctx, client, namespace, deployName)
	c.Convey("删除部署", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}
