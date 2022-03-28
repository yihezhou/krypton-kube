package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func GetK8sClient() (clientset *kubernetes.Clientset, err error){
	var config *rest.Config
	if authType == "kubeconfig" {
		//fmt.Println("use kubeconfig file to authentication")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)

	} else if authType == "password" {
		config = &rest.Config{
			Host:	clusterUrl,
			APIPath:	"/",
			Username:	userName,
			Password:	passWord,
			TLSClientConfig:	rest.TLSClientConfig{
				Insecure:true,
			},
		}
	} else if authType == "token" {
		config = &rest.Config{
			Host:	clusterUrl,
			APIPath:	"/",
			BearerToken:	bearerToken,
			TLSClientConfig:	rest.TLSClientConfig{
				Insecure:true,
			},
		}
	} else {
		log.Fatal("invaild auth_type.")
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return clientset, err
	}
	return clientset, nil
}