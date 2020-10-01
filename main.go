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
	config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-oneOCP46")
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-amqicocp45")
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config-minikube-localhost")
	//config, _ := clientcmd.LoadFromFile("/home/rgranzot/.kube/config")

	fmt.Println("===============================================================")
	fmt.Println("Playing with config")
	fmt.Println("===============================================================")

	// Kind
	fmt.Printf("Kind: %s\n", config.Kind)

	// APIVersion
	fmt.Printf("APIVersion: %s\n", config.APIVersion)

	// Preferences
	fmt.Printf("Preferences - Colors: %t\n", config.Preferences.Colors)

	// Clusters
	fmt.Printf("Clusters in config: \n")
	for _, clust := range config.Clusters {
		fmt.Printf("\tServer: %s\n", clust.Server)
		fmt.Printf("\tLocation: %s\n", clust.LocationOfOrigin)
		fmt.Printf("\tCertAuth: %s\n", clust.CertificateAuthority)
		fmt.Printf("\tSkipTLS: %t\n\n", clust.InsecureSkipTLSVerify)
	}

	// AuthInfos
	fmt.Printf("\nAuthInfos in cluster: \n")
	for _, auth := range config.AuthInfos {
		fmt.Printf("\tImpersonate: %s\n", auth.Impersonate)
		fmt.Printf("\tLocation: %s\n", auth.LocationOfOrigin)
		fmt.Printf("\tUsername: %s\n", auth.Username)
		fmt.Printf("\tAuthProvider: %s\n", auth.AuthProvider)
		fmt.Printf("\tClientKey: %s\n\n", auth.ClientKey)
	}

	// Contexts
	fmt.Printf("\nContexts in cluster: \n")
	for _, cont := range config.Contexts {
		fmt.Printf("\tCluster: %s\n", cont.Cluster)
		fmt.Printf("\tLocation: %s\n", cont.LocationOfOrigin)
		fmt.Printf("\tAuthInfo: %s\n", cont.AuthInfo)
		fmt.Printf("\tNamespace: %s\n", cont.Namespace)
		fmt.Printf("\tLocation: %s\n\n", cont.LocationOfOrigin)
	}

	// Current Context
	fmt.Printf("\nCurrentContext: %s\n", config.CurrentContext)

	// Extensions
	fmt.Printf("Extensions: %s\n", config.Extensions)

	// ====================================
	// Playing with clientConfig
	// ====================================
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{})

	fmt.Println("\n\n===============================================================")
	fmt.Println("Playing with clientConfig")
	fmt.Println("===============================================================")

	ns, override, err := clientConfig.Namespace()
	if err != nil {
		fmt.Printf("Unable to retrieve namespace from clientConfig")
		os.Exit(1)
	}

	fmt.Printf("Namespace from clientConfig : %s\n", ns)
	if override == true {
		fmt.Println("It was overwritten\n")
	}

	rc, err := clientConfig.RawConfig()
	if err != nil {
		fmt.Printf("Unable to retrieve RawConfig from clientConfig")
		os.Exit(1)
	}
	fmt.Printf("CurrentContext from Rawconfig from clientConfig : %s\n", rc.CurrentContext)

	// ====================================
	// Playing with restConfig
	// ====================================
	restConfig, _ := clientConfig.ClientConfig()

	fmt.Println("\n\n===============================================================")
	fmt.Println("Playing with restConfig")
	fmt.Println("===============================================================\n")

	fmt.Printf("Host: %s\n", restConfig.Host)
	fmt.Printf("APIPath: %s\n", restConfig.APIPath)
	fmt.Printf("ContentConfig: %s\n", restConfig.ContentConfig)
	fmt.Printf("Username: %s\n", restConfig.Username)
	fmt.Printf("Password: %s\n", restConfig.Password)
	fmt.Printf("BearerToken: %s\n", restConfig.BearerToken)
	fmt.Printf("BearerTokenFile: %s\n", restConfig.BearerTokenFile)
	fmt.Printf("Impersonate: %s\n", restConfig.Impersonate.UserName)
	//fmt.Printf("AuthProvider: %s\n", restConfig.AuthProvider.Name)
	//fmt.Printf("ExecProvider: %s\n", restConfig.ExecProvider.Command)
	fmt.Printf("TLSClientConfig: %s\n", restConfig.TLSClientConfig.ServerName)
	fmt.Printf("UserAgent: %s\n", restConfig.UserAgent)
	fmt.Printf("DisableCompression: %t\n", restConfig.DisableCompression)
	fmt.Printf("Timeout: %f\n", restConfig.Timeout.Seconds())

	// ====================================
	// Playing with kubeClient
	// ====================================
	kubeClient, _ := clientset.NewForConfig(restConfig)

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
