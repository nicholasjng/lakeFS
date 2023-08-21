package local_test

import (
	"testing"

	"github.com/treeverse/lakefs/pkg/kv/kvparams"
	"github.com/treeverse/lakefs/pkg/kv/kvtest"
	"github.com/treeverse/lakefs/pkg/kv/local"
)

func TestLocalKV(t *testing.T) {
	kvtest.DriverTest(t, local.DriverName, kvparams.Config{
		Type: local.DriverName,
		Local: &kvparams.Local{
			Path:          t.TempDir(),
			EnableLogging: true,
		},
	})
}
