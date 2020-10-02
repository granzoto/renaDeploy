package main

import (
	"fmt"
	"k8s.io/client-go/rest"
)

func restConfigFun(restConfig *rest.Config) {

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
}
