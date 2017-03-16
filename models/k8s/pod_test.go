package k8s

import (
	"testing"
)

const (
	config    = "../../conf/config.yaml"
	namespace = "default"
)

func newClient() *K8sClient {
	client, err := NewK8sClient(config)
	if err != nil {
		panic(err)
	}

	return client
}

// TestListPods
func TestListPods(t *testing.T) {

	client := newClient()

	data, err := client.ListPods(namespace)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("data:%v", data)
}
