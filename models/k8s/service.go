package k8s

import (
	"bytes"
	"html/template"

	"github.com/juju/errors"
	"github.com/zssky/log"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/pkg/api/v1"
)

func parseService(yamlfile string, data interface{}) (*v1.Service, error) {
	tpl, err := template.ParseFiles(yamlfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var svc v1.Service
	d := yaml.NewYAMLOrJSONDecoder(&b, b.Len())
	err = d.Decode(&svc)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &svc, nil
}

// ListServices - list the service
func (c *K8sClient) ListServices(namespace string) (*v1.ServiceList, error) {
	list, err := c.clientSet.Core().Services(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// CreateService - create an Service
func (c *K8sClient) CreateService(namespace, yaml string, data interface{}) (*v1.Service, error) {

	v, err := parseService(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("Service:%v", v)

	svc, err := c.clientSet.Core().Services(namespace).Create(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return svc, nil
}

// UpdateService - update Service info
func (c *K8sClient) UpdateService(namespace, yaml string, data interface{}) (*v1.Service, error) {
	v, err := parseService(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("service:%v", v)

	svc, err := c.clientSet.Core().Services(namespace).Update(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return svc, nil
}

// UpdateServiceStatus - update Service status
func (c *K8sClient) UpdateServiceStatus(namespace, yaml string, data interface{}) (*v1.Service, error) {
	v, err := parseService(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("service:%v", v)

	svc, err := c.clientSet.Core().Services(namespace).UpdateStatus(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return svc, nil
}

// DeleteService - delete an Service
func (c *K8sClient) DeleteService(namespace, name string, options *v1.DeleteOptions) error {

	err := c.clientSet.Core().Services(namespace).Delete(name, options)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// GetService - get Service detail
func (c *K8sClient) GetService(namespace, name string) (*v1.Service, error) {
	svc, err := c.clientSet.Core().Services(namespace).Get(name)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return svc, nil

}
