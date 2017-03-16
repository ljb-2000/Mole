package k8s

import (
	"bytes"
	"html/template"

	"github.com/juju/errors"
	"github.com/zssky/log"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/pkg/api/v1"
)

func parseNamespace(yamlfile string, data interface{}) (*v1.Namespace, error) {
	tpl, err := template.ParseFiles(yamlfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var ns v1.Namespace
	d := yaml.NewYAMLOrJSONDecoder(&b, b.Len())
	err = d.Decode(&ns)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &ns, nil
}

// ListNamespaces - list the Namespace
func (c *K8sClient) ListNamespaces() (*v1.NamespaceList, error) {
	list, err := c.clientSet.Core().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// CreateNamespace - create an Namespace
func (c *K8sClient) CreateNamespace(yaml string, data interface{}) (*v1.Namespace, error) {

	v, err := parseNamespace(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("Namespace:%v", v)

	ns, err := c.clientSet.Core().Namespaces().Create(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return ns, nil
}

// UpdateNamespace - update Namespace info
func (c *K8sClient) UpdateNamespace(yaml string, data interface{}) (*v1.Namespace, error) {
	v, err := parseNamespace(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("Namespace:%v", v)

	ns, err := c.clientSet.Core().Namespaces().Update(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return ns, nil
}

// UpdateNamespaceStatus - update Namespace status
func (c *K8sClient) UpdateNamespaceStatus(yaml string, data interface{}) (*v1.Namespace, error) {
	v, err := parseNamespace(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("Namespace:%v", v)

	ns, err := c.clientSet.Core().Namespaces().UpdateStatus(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return ns, nil
}

// DeleteNamespace - delete an Namespace
func (c *K8sClient) DeleteNamespace(name string, options *v1.DeleteOptions) error {

	err := c.clientSet.Core().Namespaces().Delete(name, options)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// GetNamespace - get Namespace detail
func (c *K8sClient) GetNamespace(name string) (*v1.Namespace, error) {
	ns, err := c.clientSet.Core().Namespaces().Get(name)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return ns, nil
}
