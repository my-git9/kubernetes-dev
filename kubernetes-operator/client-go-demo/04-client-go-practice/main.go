package main

import (
	"log"

	"github.com/my-git9/kubernetes-dev/client-go-demo/04-client-go-practice/pkg"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. create config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		// if no kubeconfig, by token
		inCluserConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("cat not get config")
		}
		config = inCluserConfig
	}

	// 2. create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can't create client")
	}

	// 3. create informer
	factory := informers.NewSharedInformerFactory(clientset, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()

	// 4. add informer handler
	controller := pkg.NewController(clientset, serviceInformer, ingressInformer)

	stopCh := make(chan struct{})
	// 5. start informer
	factory.Start(stopCh)
	// 等待数据同步完成后
	factory.WaitForCacheSync(stopCh)
	// 启动 controller
	controller.Run(stopCh)

}
