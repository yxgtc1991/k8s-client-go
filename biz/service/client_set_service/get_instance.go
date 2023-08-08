package client_set_service

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

var ClientSet *kubernetes.Clientset
var once sync.Once

// GetInstance 单例模式：获取clientSet
func GetInstance() *kubernetes.Clientset {
	var err error
	config := GetConf()
	once.Do(func() {
		ClientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
	})
	return ClientSet
}

// GetConf 加载配置文件
func GetConf() *rest.Config {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}
