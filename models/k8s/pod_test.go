package k8s

import (
	"testing"
)

const (
	config    = "../../template/config.yaml"
	namespace = "default"
	podyaml   = "../../template/vttablet-pod-template.yaml"
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

// TestCreatePod
func TestCreatePod(t *testing.T) {
	client := newClient()

	v := VttabletPodData{
		Uid:               "488484",
		Keyspace:          "test_keyspace",
		Shard_label:       "xx-80",
		Alias:             "test-0000001100",
		Vitess_image:      "192.168.212.19/vitesss/lite:v1.0.1",
		Port:              15002,
		Grpc_port:         15999,
		Tablet_type:       "replica",
		Backup_flags:      "-backup_storage_implementation file -file_backup_storage_root '/vt/vtdataroot/backup/'",
		Vtdataroot_volume: "emptyDir : {}",
		Shard:             "-80",
	}

	pod, err := client.CreatePod(namespace, podyaml, v)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("pod:%v", pod)
}

// TestDeletePod
func TestDeletePod(t *testing.T) {
	name := "vttablet-488484"

	client := newClient()
	err := client.DeletePod(namespace, name, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("delete pod success")
}

// TestGetPod
func TestGetPod(t *testing.T) {
	name := "vttablet-488484"

	client := newClient()
	pod, err := client.GetPod(namespace, name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("pod:%v", pod)
}
