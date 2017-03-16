package k8s

import (
	"k8s.io/client-go/pkg/api/v1"

	"github.com/juju/errors"
)

// ListPods - list the pods
func (c *K8sClient) ListPods(namespace string) (*v1.PodList, error) {
	pods, err := c.clientSet.Core().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pods, nil
}
