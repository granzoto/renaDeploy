package main

import (
	"fmt"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	// Create Interconnect Pod
	// Using a yaml
	// ====================================
	err := CreatePodFromYaml(kubeClient, simpleic, "bundlenato46")
	if err != nil {
		fmt.Println("Unable to deploy IC pod from yaml",  err)
	}

	// ====================================
	// Create Interconnect Deployment
	// Using Jurutypes
	// ====================================
	fmt.Println("Starting deployGo")
	deployGo(kubeClient)

	//// Check the Deployment replicas - Note that I deployed IC froma yaml file manually, ATM
	//deps, err := kubeClient.AppsV1().Deployments("bundlenato46").List(context.TODO(), v1.ListOptions{})
	//if err != nil {
	//	fmt.Println("Unable to retrieve deployments")
	//	os.Exit(1)
	//}
	//
	//for _, dep := range deps.Items {
	//	fmt.Printf("Deployment %s\n", dep.Name)
	//	fmt.Printf("Desired Replicas %d\n", dep.Status.Replicas)
	//	fmt.Printf("Available Replicas %d\n", dep.Status.AvailableReplicas)
	//	// Check what else we need to get from the deployment
	//}
	//
	//svcs, err := kubeClient.CoreV1().Services("bundlenato46").List(context.TODO(), v1.ListOptions{})
	//if err != nil {
	//	fmt.Println("Unable to retrieve services")
	//	os.Exit(1)
	//}
	//
	//for _, svc := range svcs.Items {
	//	fmt.Printf("Service %s\n", svc.Name)
	//	fmt.Printf("Service Ports : \n")
	//	for _, port := range svc.Spec.Ports {
	//		fmt.Printf("\tPort: %d - Name: %s - Proto:%s\n", port.Port, port.Name, port.Protocol)
	//	}
	//	//fmt.Printf("Available Replicas %d\n", dep.Status.AvailableReplicas)
	//	// Check what else we need to get from the deployment
	//}
}