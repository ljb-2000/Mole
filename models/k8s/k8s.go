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
	Uid              string `json:"Uid"`
	Keyspace         string `json:"Keyspace"`
	ShardLabel       string `json:"ShardLable"`
	Alias            string `json:"Alias"`
	VitessImage      string `json:"VitessImage"`
	Port             int    `json:"Port"`
	GrpcPort         int    `json:"GrpcPort"`
	TabletType       string `json:"TabletType"`
	BackupFlags      string `json:"BackupFlags"`
	VtdatarootVolume string `json:"VtdatarootVolume"`
	Shard            string `json:"Shard"`
	TabletSubdir     string `json:"TabletSubdir"`
}

// VtgateServiceData
type VtgateServiceData struct {
	Type            string `json:"Type"`
	Cell            string `json:"Cell"`
	ServiceType     string `json"ServiceType"`
	MysqlServerPort int    `json:"MysqlServerPort"`
}

// VtgateRCData
type VtgateRCData struct {
	Type             string `json:"Type"`
	Cell             string `json:"Cell"`
	VitessImage      string `json:"VitessImage"`
	Replicas         int    `json:"Replicas"`
	MysqlServerPort  int    `json:"MysqlServerPort"`
	BackupFlags      string `json:"BackupFlags"`
	TestFlags        string `json:"TestFlags"`
	VtdatarootVolume string `json:"VtdatarootVolume"`
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
