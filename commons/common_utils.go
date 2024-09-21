package common_client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func Create_clientset() (*kubernetes.Clientset, error) {
	config, config_err := clientcmd.BuildConfigFromFlags("", os.Args[1])
	if config_err != nil {
		fmt.Println("Error in loading kube config file", config_err)
	}
	clientset, client_err := kubernetes.NewForConfig(config)
	if client_err != nil {
		fmt.Println("Error in loading the clientset", client_err)
	}
	return clientset, client_err
}

func Create_namespace_client(clientset *kubernetes.Clientset) apiv1.NamespaceInterface {
	return clientset.CoreV1().Namespaces()
}

func Create_node_client(clientset *kubernetes.Clientset) apiv1.NodeInterface {
	return clientset.CoreV1().Nodes()
}

func Create_pod_client(clientset *kubernetes.Clientset, namespace string) apiv1.PodInterface {
	return clientset.CoreV1().Pods(namespace)
}
