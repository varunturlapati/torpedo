package schedops

import (
	"github.com/portworx/torpedo/drivers/node"
	"github.com/Sirupsen/logrus"
	"github.com/portworx/torpedo/pkg/errors"
)

// Driver is the interface for portworx operations under various schedulers
type Driver interface {
	// DisableOnNode disabled portworx on given node
	DisableOnNode(n node.Node) error
	// ValidateOnNode validates portworx on given node (from scheduler perspective)
	ValidateOnNode(n node.Node) error
	// EnableOnNode enabled portworx on given node
	EnableOnNode(n node.Node) error
}

var (
	schedOpsRegistry = make(map[string]Driver)
)


// Register registers the given portworx scheduler operator
func Register(name string, d Driver) error {
	logrus.Infof("Registering portworx scheduler operator: %v", name)
	schedOpsRegistry[name] = d
	return nil
}

// Get a driver to perform portworx operations for the given scheduler
func Get(name string) (Driver, error) {
	d, ok := schedOpsRegistry[name]
	if ok {
		return d, nil
	}

	return nil, &errors.ErrNotFound{
		ID:   name,
		Type: "Portworx Scheduler Operator",
	}
}