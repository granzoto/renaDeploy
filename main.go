package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

//const (
//	NS = "openshift"
//)

func main() {

	// ====================================
	// Playing with config
	// ====================================
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-amqicocp45")
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-minikube-localhost")
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config")
	config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-oneOCP46")
    //configFun(config)

	// ====================================
	// Playing with clientConfig
	// ====================================
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{})
    //clientConfigsFun(clientConfig)

	// ====================================
	// Playing with restConfig
	// ====================================
	restConfig, _ := clientConfig.ClientConfig()
    //restConfigFun(restConfig)

	// ====================================
	// Playing with kubeClient
	// ====================================
	kubeClient, _ := clientset.NewForConfig(restConfig)
	//listResources(kubeClient)

	// ====================================
	// Create Interconnect Deployment
	// Using a yaml
	// ====================================


	// Check the Deployment replicas - Note that I deployed IC froma yaml file manually, ATM
	deps, err := kubeClient.AppsV1().Deployments("bundlenato46").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to retrieve deployments")
		os.Exit(1)
	}

	for _, dep := range deps.Items {
		fmt.Printf("Deployment %s\n", dep.Name)
		fmt.Printf("Desired Replicas %d\n", dep.Status.Replicas)
		fmt.Printf("Available Replicas %d\n", dep.Status.AvailableReplicas)
		// Check what else we need to get from the deployment
	}

	svcs, err := kubeClient.CoreV1().Services("bundlenato46").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to retrieve services")
		os.Exit(1)
	}

	for _, svc := range svcs.Items {
		fmt.Printf("Service %s\n", svc.Name)
		fmt.Printf("Service Status %s\n", svc.Status)
		// How to get the service status correctly
		
		//fmt.Printf("Available Replicas %d\n", dep.Status.AvailableReplicas)
		// Check what else we need to get from the deployment
	}
}
