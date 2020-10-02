package main

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd/api"
)

func configFun(config *api.Config) {

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
}
