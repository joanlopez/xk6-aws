package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/aws", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// aws is the exported type
		aws *AWS
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu:  vu,
		aws: &AWS{vu: vu},
	}
}

// AWS is the type for our custom API.
type AWS struct {
	vu modules.VU // provides methods for accessing internal k6 objects

	s3 s3.Client
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.aws,
	}
}
