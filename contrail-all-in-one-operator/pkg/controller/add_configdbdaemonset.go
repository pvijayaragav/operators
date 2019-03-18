package controller

import (
	"github.com/contrail-operators/contrail-all-in-one-operator/pkg/controller/configdbdaemonset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, configdbdaemonset.Add)
}
