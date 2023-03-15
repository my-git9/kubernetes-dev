package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

func main() {
	var kubeConfig *string

	ctx := context.Background()

	// absolute kubeconfig path flag
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute the kubeconfig path")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute the kubeconfig path")
	}
	namespace := flag.String("namespace", "defalut", "specify namespace")
	// 解析参数
	flag.Parse()
	// clientcmd.BuildConfigFromFlags： k8s 命令行解析工具
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		// klog: k8s的日志工具
		klog.Fatal(err)
	}

	//
	clientSet, err := kubernetes.NewForConfig(config)

	// get ns
	namespacelist, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		// klog: k8s的日志工具
		klog.Fatal(err)
	}

	namespaces := namespacelist.Items
	for _, namespace := range namespaces {
		fmt.Println("name ===> "+ namespace.Name + " ===> Status: "+ string(namespace.Status.Phase))
	}

	// get po
	podlist, err := clientSet.CoreV1().Pods(*namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		// klog: k8s的日志工具
		klog.Fatal(err)
	}
	pods := podlist.Items
	for _, pod := range pods{
		fmt.Println(pod.Name)
	}
}
