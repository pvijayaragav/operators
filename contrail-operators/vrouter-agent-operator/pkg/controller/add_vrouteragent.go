package controller

import (
	"github.com/contrail-operators/vrouter-agent-operator/pkg/controller/vrouteragent"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, vrouteragent.Add)
}
