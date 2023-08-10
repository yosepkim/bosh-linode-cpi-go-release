package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	"bosh-linode-cpi/api"
	instance "bosh-linode-cpi/linode/instance_service"
)

type DeleteVM struct {
	vmService instance.Service
}

func NewDeleteVM(
	vmService instance.Service,
) DeleteVM {
	return DeleteVM{
		vmService: vmService,
	}
}

func (dv DeleteVM) Run(vmCID VMCID) (interface{}, error) {
	// Delete the VM
	if err := dv.vmService.Delete(string(vmCID)); err != nil {
		if _, ok := err.(api.CloudError); ok {
			return nil, err
		}
		return nil, bosherr.WrapErrorf(err, "Deleting vm '%s'", vmCID)
	}

	return nil, nil
}
