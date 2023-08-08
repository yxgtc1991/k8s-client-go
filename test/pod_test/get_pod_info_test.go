package pod_test

import (
	"context"
	"demo/k8s-client-go/biz/service/client_set_service"
	"demo/k8s-client-go/biz/service/pod_service"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPodInfo(t *testing.T) {
	ctx := context.TODO()
	client := client_set_service.GetInstance()
	namespace := "default"
	err := pod_service.NewPodService().GetPodInfo(ctx, client, namespace)
	c.Convey("获取Pod详情", t, func() {
		c.So(err, c.ShouldEqual, nil)
	})
}
