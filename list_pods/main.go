package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func get_kube_config_path() string {
	var kube_config_path string
	home_dir := homedir.HomeDir()

	if _, err := os.Stat(home_dir + "/.kube/config"); err == nil {
		kube_config_path = home_dir + "/.kube/config"
	} else {
		fmt.Println("Enter kubernetes config directory: ")
		fmt.Scanf("%s", kube_config_path)
	}

	return kube_config_path
}

func main() {
	// Get kube config path
	kube_config_path := get_kube_config_path()

	fmt.Println(kube_config_path)

	// Build configuration from config file
	config, err := clientcmd.BuildConfigFromFlags("", kube_config_path)
	if err != nil {
		panic(err)
	}

	// Create clientser
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Create pods
	pod, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Print list pods
	for _, pod := range pod.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
	}
}
