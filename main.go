package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	klient "github.com/slayer321/crdcreation/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Building config from flags failed, %s, trying to build inclusterconfig", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Printf("error %s building inclusterconfig", err.Error())
		}
	}

	clientset, err := klient.NewForConfig(config)
	if err != nil {
		fmt.Println("Not able to get the clientset")
	}

	//clientset.SachinmauryaV1alpha1().HADeployments("").List(context.Background(), metav1.ListOptions{})
	//fmt.Println(clientset)

	item, err := clientset.SachinmauryaV1alpha1().HADeployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("not able to list the item: %s", err.Error())
	}

	fmt.Println("Len if HAdepployment is", len(item.Items))
}
