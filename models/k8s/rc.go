package k8s

import (
	"bytes"
	"html/template"

	"github.com/juju/errors"
	"github.com/zssky/log"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/pkg/api/v1"
)

func parseRC(yamlfile string, data interface{}) (*v1.ReplicationController, error) {
	tpl, err := template.ParseFiles(yamlfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var rc v1.ReplicationController
	d := yaml.NewYAMLOrJSONDecoder(&b, b.Len())
	err = d.Decode(&rc)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &rc, nil
}

// ListRC - list the ReplicationController
func (c *K8sClient) ListRC(namespace string) (*v1.ReplicationControllerList, error) {
	list, err := c.clientSet.Core().ReplicationControllers(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// CreateRC - create an ReplicationController
func (c *K8sClient) CreateRC(namespace, yaml string, data interface{}) (*v1.ReplicationController, error) {

	v, err := parseRC(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("ReplicationController:%v", v)

	rc, err := c.clientSet.Core().ReplicationControllers(namespace).Create(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return rc, nil
}

// UpdateRC - update ReplicationController info
func (c *K8sClient) UpdateRC(namespace, yaml string, data interface{}) (*v1.ReplicationController, error) {
	v, err := parseRC(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("ReplicationController:%v", v)

	rc, err := c.clientSet.Core().ReplicationControllers(namespace).Update(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return rc, nil
}

// UpdateRCStatus - update ReplicationController status
func (c *K8sClient) UpdateRCStatus(namespace, yaml string, data interface{}) (*v1.ReplicationController, error) {
	v, err := parseRC(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("ReplicationController:%v", v)

	rc, err := c.clientSet.Core().ReplicationControllers(namespace).UpdateStatus(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return rc, nil
}

// DeleteRC - delete an ReplicationController
func (c *K8sClient) DeleteRC(namespace, name string, options *v1.DeleteOptions) error {

	err := c.clientSet.Core().ReplicationControllers(namespace).Delete(name, options)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// GetRC- get ReplicationController detail
func (c *K8sClient) GetRC(namespace, name string) (*v1.ReplicationController, error) {
	rc, err := c.clientSet.Core().ReplicationControllers(namespace).Get(name)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return rc, nil
}
