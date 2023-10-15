package main

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main()  {
	// 1. create config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/xin/.kube/config")
	if err != nil{
		panic(err)
	}

	// 2. create client
	clientset, err := kubernetes.NewForConfig(config)

	// 3. create informer
	//factory := informers.NewSharedInformerFactory(clientset, 0)
	// 指定命名空间
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))
	informer := factory.Core().V1().Pods().Informer()

	// add workqueue
	rateLimitingQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller")

	// 4. add event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("ADD Event")
			key, err := cache.MetaNamespaceIndexFunc(obj)
			if err != nil {
				fmt.Println("cat't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update Event")
			key, err := cache.MetaNamespaceIndexFunc(newObj)
			if err != nil {
				fmt.Println("cat't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete Event")
			key, err := cache.MetaNamespaceIndexFunc(obj)
			if err != nil {
				fmt.Println("cat't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
	})

	// 5. start informer
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<- stopCh
}
