package k8s

import (
	"testing"
)

const (
	server = "http://192.168.180.101:8080"
)

// TestGetPods
func TestGetPods(t *testing.T) {
	namespace := "default"

	data, err := GetPods(server, namespace)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("data:%v", data)
}
