package main

import (
	"context"
	"fmt"
	common_client "sample/commons"
	namespaces "sample/namespaces"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Custom struct {
	Namespace string
	Pod       string
	NodeName  string
}

func Custom_struct(clientset *kubernetes.Clientset) []Custom {
	Customs := []Custom{}
	namespace_list := namespaces.List_Namespace(common_client.Create_namespace_client(clientset))
	for _, namespace := range namespace_list.Items {
		pod_list, pod_err := common_client.Create_pod_client(clientset, namespace.Name).List(context.Background(), v1.ListOptions{})
		if pod_err != nil {
			fmt.Println("Error in pod list", pod_err)
		}

		if len(pod_list.Items) == 0 {
			continue
		}
		for _, pod := range pod_list.Items {
			Customs = append(Customs, Custom{
				Namespace: namespace.Name,
				Pod:       pod.Name,
				NodeName:  pod.Spec.NodeName,
			})
		}
	}
	return Customs
}

func main() {
	fmt.Println("Starting the script")

	clientset, error := common_client.Create_clientset()
	if error != nil {
		fmt.Println("Error in clientset", clientset)
	}

	Custom_output := Custom_struct(clientset)
	// fmt.Println(Custom_output)
	
	for _, custom_struct := range Custom_output {
		// var deployments string
		// deployment_list := namespaces.List_Deployment(clientset, custom_struct.Namespace)
		// // fmt.Println(deployment_list)
		// fmt.Println(deployment_list.Items[0].Spec)
		// if len(deployment_list.Items) > 0 {
		// 	for _, val := range namespaces.Loop_deployment(deployment_list) {
		// 		deployments = deployments + val
		// 	}
		fmt.Printf("Namespaces %v contains pods %v and is scheduled on %v .\n", custom_struct.Namespace, custom_struct.Pod, custom_struct.NodeName)
		
	}

}
