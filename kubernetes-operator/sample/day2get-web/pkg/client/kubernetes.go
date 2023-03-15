package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
)

// create clientset
func GetK8sClientSet() (*kubernetes.Clientset, error) {
	config, err := GetRestConfig()
	if err != nil {
		return nil, err
	}
	// 创建 clientset （客户集）
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	return clientSet, nil
}

// get config
func GetRestConfig() (config *rest.Config, err error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute the kubeconfig path")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute the kubeconfig path")
	}
	flag.Parse()

	// clientcmd.BuildConfigFromFlags： k8s 命令行解析工具
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// klog: k8s的日志工具
		klog.Fatal(err)
		return
	}
	return config, nil
}
