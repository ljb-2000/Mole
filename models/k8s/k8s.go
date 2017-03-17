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
	Uid               string `json:"Uid"`
	Keyspace          string `json:"Keyspace"`
	Shard_label       string `json:"ShardLable"`
	Alias             string `json:"Alias"`
	Vitess_image      string `json:"VitessImage"`
	Port              int    `json:"Port"`
	Grpc_port         int    `json:"GrpcPort"`
	Tablet_type       string `json:"TabletType"`
	Backup_flags      string `json:"BackupFlags"`
	Vtdataroot_volume string `json:"VtdatarootVolume"`
	Shard             string `json:"Shard"`
}

// VtgateServiceData
type VtgateServiceData struct {
	Cell            string `json:"Cell"`
	MysqlServerPort int    `json:"MysqlServerPort"`
}

// VtgateRCData
type VtgateRCData struct {
	Cell            string `json:"Cell"`
	VitessImage     string `json:"VitessImage"`
	Replicas        int    `json:"Replicas"`
	MysqlServerPort int    `json:"MysqlServerPort"`
}

// NamespaceData
type NamespaceData struct {
	Name   string            `json:"Name"`
	Labels map[string]string `json:"Labels"`
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
