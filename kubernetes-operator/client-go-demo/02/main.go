package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main()  {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/xin/.kube/config")
	if err != nil{
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	deploymentsClient, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil{
		panic(err)
	} else {
		deploymentlist := deploymentsClient.Items
		for _, deployment := range deploymentlist{
			println(deployment.Name)
		}
	}
}


