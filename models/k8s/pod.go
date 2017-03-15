package k8s

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/kubernetes/kubernetes/pkg/api"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

// GetPods - get pods list
func GetPods(server string, namespace string) (*api.PodList, error) {
	url := fmt.Sprintf(pods, namespace)

	data, _, err := http.SimpleGet(url, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var podList *api.PodList
	if err := tc.DecodeJSON(data, podList); err != nil {
		return nil, errors.Trace(err)
	}

	return podList, nil
}
