package informer_test

import (
	"demo/k8s-client-go/biz/service/client_set_service"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"log"
	"testing"
)

func TestSharedInformer(t *testing.T) {
	client := client_set_service.GetInstance()
	// 1、初始化 informer factory
	factory := informers.NewSharedInformerFactoryWithOptions(
		client, 0, informers.WithNamespace("default"))
	// 2、初始化 pod informer
	informer := factory.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod Updated to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod Deleted from Store: %s", mObj.GetName())
		},
	})
	stopCh := make(chan struct{})
	// 3、启动 informer factory
	factory.Start(stopCh)
	// 4、等待list操作获取到的对象都同步到informer本地缓存Indexer中
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
