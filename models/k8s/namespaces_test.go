package k8s

import (
	"testing"
)

const (
	namespaceyaml = "../../template/namespace-template.yaml"
)

// TestListNamespaces
func TestListNamespaces(t *testing.T) {
	client := newClient()

	data, err := client.ListNamespaces()
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("data:%v", data)
}

// TestCreateNamespace
func TestCreateNamespace(t *testing.T) {
	client := newClient()

	v := NamespaceData{
		Name: "test-ns",
		Labels: map[string]string{
			"name": "test-ns",
			"app":  "test",
		},
	}

	ns, err := client.CreateNamespace(namespaceyaml, v)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("Namespace:%v", ns)
}

// TestDeleteNamespace
func TestDeleteNamespace(t *testing.T) {
	name := "test-ns"

	client := newClient()
	err := client.DeleteNamespace(name, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("delete Namespace success")
}

// TestGetNamespace
func TestGetNamespace(t *testing.T) {
	name := "test-ns"

	client := newClient()
	ns, err := client.GetNamespace(name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("Namespace:%v", ns)
}
