package k8s

import (
	"testing"
)

const (
	serviceyaml = "../../template/vtgate-service-template.yaml"
)

// TestListServices
func TestListServices(t *testing.T) {
	client := newClient()

	data, err := client.ListServices(namespace)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("data:%v", data)
}

// TestCreateService
func TestCreateService(t *testing.T) {
	client := newClient()

	v := VtgateServiceData{
		Cell:            "test11",
		MysqlServerPort: 15000,
	}

	svc, err := client.CreateService(namespace, serviceyaml, v)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("service:%v", svc)
}

// TestDeleteService
func TestDeleteService(t *testing.T) {
	name := "vtgate-test11"

	client := newClient()
	err := client.DeleteService(namespace, name, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("delete Service success")
}

// TestGetService
func TestGetService(t *testing.T) {
	name := "vtgate-test11"

	client := newClient()
	svc, err := client.GetService(namespace, name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("Service:%v", svc)
}
