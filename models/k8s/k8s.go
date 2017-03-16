package k8s

import (
	"github.com/juju/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	kubeConfig string
	clientSet  *kubernetes.Clientset
}

// NewK8sClient - create an new k8s client
func NewK8sClient(cfgfile string) (*K8sClient, error) {

	config, err := clientcmd.BuildConfigFromFlags("", cfgfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Create the ClientSet
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &K8sClient{
		kubeConfig: cfgfile,
		clientSet:  client,
	}, nil
}
