package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func listResources(kubeClient *kubernetes.Clientset) {

	fmt.Println("\n\n===============================================================")
	fmt.Println("Playing with kubeClient")
	fmt.Println("===============================================================\n")

	//====================================
	// Listing namespaces
	//====================================
	nsList, err := kubeClient.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, ns := range nsList.Items {
		fmt.Printf("Namespace: %s\n", ns.Name)
	}

	//====================================
	// Listing pods
	//====================================
	podList, err := kubeClient.CoreV1().Pods("openshift-monitoring").List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		fmt.Printf("Pod: %s", pod.Name)
	}

	//====================================
	// Listing Containers from Pods
	//====================================
	podList, err = kubeClient.CoreV1().Pods("openshift-monitoring").List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		fmt.Printf("\nPod: %s\n", pod.Name)
		for _, cont := range pod.Spec.Containers {
			fmt.Printf("\t- Container: %s\n", cont.Name)
			fmt.Printf("\t\t- Image: %s\n", cont.Image)
			fmt.Printf("\t\t- Ports:\n")
			for _, port := range cont.Ports {
				fmt.Printf("\t\t\t - Port Name: %s\n", port.Name)
				fmt.Printf("\t\t\t - Cont Port: %d\n", port.ContainerPort)
				fmt.Printf("\t\t\t - Host IP: %s\n", port.HostIP)
				fmt.Printf("\t\t\t - Host Port: %d\n", port.HostPort)
				fmt.Printf("\t\t\t - Protocol: %s\n\n", port.Protocol)
			}

		}
	}
}
