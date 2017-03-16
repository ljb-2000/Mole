package k8s

import (
	"testing"
)

const (
	rcyaml = "../../template/vtgate-controller-template.yaml"
)

// TestListRC
func TestListRC(t *testing.T) {
	client := newClient()

	data, err := client.ListRC(namespace)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("data:%v", data)
}

// TestCreateRC
func TestCreateRC(t *testing.T) {
	client := newClient()

	v := VtgateRCData{
		Cell:            "test11",
		VitessImage:     "192.168.212.19/vitesss/lite:v1.0.1",
		Replicas:        2,
		MysqlServerPort: 3306,
	}

	rc, err := client.CreateRC(namespace, rcyaml, v)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("rc:%v", rc)
}

// TestDeleteRC
func TestDeleteRC(t *testing.T) {
	name := "vtgate-test11"

	client := newClient()
	err := client.DeleteRC(namespace, name, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("delete rc success")
}

// TestGetRC
func TestGetRC(t *testing.T) {
	name := "vtgate-test11"

	client := newClient()
	rc, err := client.GetRC(namespace, name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("rc:%v", rc)
}
