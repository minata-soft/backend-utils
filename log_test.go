package backend_utils_test

import (
	"testing"

	backend_utils "github.com/minata-soft/backend-utils"
)

func TestDebugMessage(t *testing.T) {
	a := "test"
	b := "deux"
	backend_utils.Debug.Debug("test de ma function %s %s", a, b)
}

func TestErrorMessage(t *testing.T) {
	a := "test"
	b := "deux"
	backend_utils.Debug.Error("test de ma function %s %s", a, b)
}

func TestInfo(t *testing.T) {
	a := "test"
	b := "deux"
	backend_utils.Debug.Info("test de ma function %s %s", a, b)
}
