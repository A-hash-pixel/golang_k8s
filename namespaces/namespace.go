package namespaces

import (
	"context"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	appsv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// List namespaces
func List_Namespace(namespace_client apiv1.NamespaceInterface) *appsv1.NamespaceList {
	namespace_list, err := namespace_client.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return namespace_list
}

// List nodes
func List_Nodes(node_client apiv1.NodeInterface) *appsv1.NodeList {
	node_list, node_error := node_client.List(context.Background(), metav1.ListOptions{})
	if node_error != nil {
		fmt.Println("Error in listing node", node_error.Error())
	}
	return node_list
}

// List Deployments
func List_Deployment(clientset *kubernetes.Clientset, namepsace string) *v1.DeploymentList {
	deployment_list, err := clientset.AppsV1().Deployments(namepsace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error in listing deployment", err)
	}
	return deployment_list
}

// Loop through deployment list
func Loop_deployment(deployment_list *v1.DeploymentList) []string {
	deployments := []string{}
	for _, val := range deployment_list.Items {
		deployments = append(deployments, val.Name)
	}

	return deployments
}