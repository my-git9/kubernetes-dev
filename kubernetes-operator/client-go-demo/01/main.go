package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main()  {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/xin/.kube/config")
	if err != nil{
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "api"

	// client
	restclient, err := rest.RESTClientFor(config)
	if err != nil{
		panic(err)
	}

	// get data
	pod := v1.Pod{}
	err = restclient.Get().Namespace("default").Resource("pods").Name("details-v1-5f6994d866-jf2dh").Do(context.TODO()).Into(&pod)
	if err != nil{
		panic(err)
	} else {
		println(pod.Status.PodIP)
	}
}
