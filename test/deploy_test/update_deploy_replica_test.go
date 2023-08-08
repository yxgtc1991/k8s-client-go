package deploy

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/deploy_service"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUpdateDeployReplicaTest(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	deployName := "httpd"
	replica := int32(5)
	err := deploy_service.NewDeployService().UpdateDeployReplica(ctx, client, namespace, deployName, replica)
	c.Convey("更新部署副本数", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}

func TestUpdateDeployReplicaTest2(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	deployName := "httpd"
	replica := int32(3)
	err := deploy_service.NewDeployService().UpdateDeployReplica2(ctx, client, namespace, deployName, replica)
	c.Convey("更新部署副本数", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}
