package main

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func clientConfigsFun(clientConfig clientcmd.ClientConfig) {

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
}
