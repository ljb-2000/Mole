package vtctl

import (
	"fmt"
	"time"

	"github.com/juju/errors"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

const (
	api_prefix   = "/api/"
	keyspaces    = "keyspaces/"
	vtctl        = "vtctl/"
	shards       = "shards/"
	tablets      = "tablets/"
	schema_apply = "schema/apply"

	deadline    = time.Second * 30
	dialtimeout = time.Second * 5
)

// VtctlResponse - vtctl api response data
type VtctlResponse struct {
	Error  string `json:"Error"`
	Output string `json:"output"`
}

// Tablet
type Tablet struct {
	Cell string `json:"cell"`
	Uid  string `json:"uid"`
}

// KeyspacesList - get Keyspaces list
func KeyspacesList(server string) ([]string, error) {

	url := fmt.Sprintf("%v%v%v", server, api_prefix, keyspaces)
	data, _, err := http.SimpleGet(url, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([]string, 0)

	err = tc.DecodeJSON(data, list)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// Vtctl - run vtctl command
// @param1 out:string
// @param2 err:error
func Vtctl(server string, v interface{}) (*VtctlResponse, error) {
	url := fmt.Sprintf("%v%v%v", server, api_prefix, vtctl)
	data, _, err := http.PostJSON(url, v, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var resp VtctlResponse
	err = tc.DecodeJSON(data, resp)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &resp, nil
}

// Shards - get shards list
func Shards(server string, keyspace string) ([]string, error) {
	url := fmt.Sprintf("%v%v%v%v", server, api_prefix, shards, keyspace)
	data, _, err := http.SimpleGet(url, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([]string, 0)
	err = tc.DecodeJSON(data, list)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// Tablets
func Tablets(server string, shard string) ([]Tablet, error) {
	url := fmt.Sprintf("%v%v%v", server, api_prefix, tablets)
	data := []byte(fmt.Sprintf("shard=%v", shard))
	data, _, err := http.PostForm(url, data, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([]Tablet, 0)
	err = tc.DecodeJSON(data, list)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return list, nil
}

// SchemaApply
func SchemaApply(server string, keyspace string, sql string) (*VtctlResponse, error) {
	url := fmt.Sprintf("%v%v%v", server, api_prefix, schema_apply)
	v := struct {
		Keyspace string
		SQL      string
	}{
		keyspace,
		sql,
	}

	data, _, err := http.PostJSON(url, v, deadline, dialtimeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var resp VtctlResponse
	err = tc.DecodeJSON(data, resp)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &resp, nil
}
