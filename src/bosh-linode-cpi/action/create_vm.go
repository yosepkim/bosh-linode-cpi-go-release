package action

import (
	"bosh-linode-cpi/api"
	instance "bosh-linode-cpi/linode/instance_service"
	"bosh-linode-cpi/registry"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type CreateVM struct {
	vmService       instance.Service
	registryOptions registry.ClientOptions
}

func NewCreateVM(
	vmService instance.Service,
) CreateVM {
	return CreateVM{vmService: vmService}
}

func (cv CreateVM) Run(agentID string, stemcellCID StemcellCID, cloudProps VMCloudProperties) (VMCID, error) {
	var err error
	var vm string
	if err = cloudProps.Validate(); err != nil {
		return "", bosherr.WrapError(err, "Creating VM")
	}

	vmProps := &instance.Properties{
		LinodeType: cloudProps.LinodeType,
		Region:     cloudProps.Region,
		Tags:       cloudProps.Tags,
	}

	// Create VM
	vm, err = cv.vmService.Create(vmProps, cv.registryOptions.EndpointWithCredentials())
	if err != nil {
		if _, ok := err.(api.CloudError); ok {
			return "", err
		}
		return "", bosherr.WrapError(err, "Creating VM")
	}
	vmCID := VMCID(vm)

	// If any of the below code fails, we must delete the created vm
	defer func() {
		if err != nil {
			cv.vmService.CleanUp(vm)
		}
	}()

	return vmCID, nil
}
