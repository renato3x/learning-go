package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
  if runtime.GOARCH != "amd64" {
    t.Skip("Skipping test on non-amd64 architecture")
  }

  t.Fail()
}
