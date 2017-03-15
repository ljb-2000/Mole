package k8s

import (
	"flag"
	"fmt"
	"time"

	"github.com/juju/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	kubeConfig string
	clientSet  *kubernetes.Clientset
}

// NewK8sClient - create an new k8s client
func NewK8sClient(config string) (*K8sClient, error) {

	config, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Create the ClientSet
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &K8sClient{
		kubeconfig: config,
		clientSet:  client,
	}, nil
}

// ListPods - list the pods
func (c *K8sClient) ListPods() {
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
}
