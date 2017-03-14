package k8s

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
