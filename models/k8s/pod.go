package k8s

import (
	"bytes"
	"html/template"

	"github.com/juju/errors"
	"github.com/zssky/log"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/pkg/api/v1"
)

func parsePod(yamlfile string, data interface{}) (*v1.Pod, error) {
	tpl, err := template.ParseFiles(yamlfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var pod v1.Pod
	d := yaml.NewYAMLOrJSONDecoder(&b, b.Len())
	err = d.Decode(&pod)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &pod, nil
}

// ListPods - list the pods
func (c *K8sClient) ListPods(namespace string) (*v1.PodList, error) {
	pods, err := c.clientSet.Core().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pods, nil
}

// CreatePod - create an pod
func (c *K8sClient) CreatePod(namespace, yaml string, data interface{}) (*v1.Pod, error) {

	v, err := parsePod(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("pod:%v", v)

	pod, err := c.clientSet.Core().Pods(namespace).Create(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pod, nil
}

// UpdatePod - update pod info
func (c *K8sClient) UpdatePod(namespace, yaml string, data interface{}) (*v1.Pod, error) {
	v, err := parsePod(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("pod:%v", v)

	pod, err := c.clientSet.Core().Pods(namespace).Update(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pod, nil
}

// UpdatePodStatus - update pod status
func (c *K8sClient) UpdatePodStatus(namespace, yaml string, data interface{}) (*v1.Pod, error) {
	v, err := parsePod(yaml, data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Debugf("pod:%v", v)

	pod, err := c.clientSet.Core().Pods(namespace).UpdateStatus(v)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pod, nil
}

// DeletePod - delete an pod
func (c *K8sClient) DeletePod(namespace, name string, options *v1.DeleteOptions) error {

	err := c.clientSet.Core().Pods(namespace).Delete(name, options)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// GetPod - get pod detail
func (c *K8sClient) GetPod(namespace, name string) (*v1.Pod, error) {
	pod, err := c.clientSet.Core().Pods(namespace).Get(name)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pod, nil
}
