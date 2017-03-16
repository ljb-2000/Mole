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

// VttabletPodData
type VttabletPodData struct {
	Uid               string
	Keyspace          string
	Shard_label       string
	Alias             string
	Vitess_image      string
	Port              int
	Grpc_port         int
	Tablet_type       string
	Backup_flags      string
	Vtdataroot_volume string
	Shard             string
}

// VtgateServiceData
type VtgateServiceData struct {
	Cell            string
	MysqlServerPort int
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
