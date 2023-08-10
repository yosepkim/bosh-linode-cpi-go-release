package instance

import (
	"context"
	"strconv"

	"github.com/linode/linodego"

	"bosh-linode-cpi/api"
)

func (i LinodeInstanceService) Create(vmProps *Properties, registryEndpoint string) (string, error) {
	boolTrue := true
	swapSize := 0
	instanceCreateOptions := linodego.InstanceCreateOptions{
		Region:   vmProps.Region,
		Type:     vmProps.LinodeType,
		Booted:   &boolTrue,
		SwapSize: &swapSize,
		Tags:     vmProps.Tags,
		// TODO: label
	}

	i.logger.Debug(linodeInstanceServiceLogTag, "Creating Linode Instance with params: %v", instanceCreateOptions)
	instance, err := i.linodeClient.CreateInstance(context.Background(), instanceCreateOptions)
	if err != nil {
		i.logger.Debug(linodeInstanceServiceLogTag, "Failed to create Linode Instance: %v", err)
		return "", api.NewVMCreationFailedError(err.Error())
	}

	return strconv.Itoa(instance.ID), nil
}

func (i LinodeInstanceService) CleanUp(id string) {
	if linodeId, err := strconv.Atoi(id); err != nil {
		if err := i.linodeClient.DeleteInstance(context.Background(), linodeId); err != nil {
			i.logger.Debug(linodeInstanceServiceLogTag, "Failed cleaning up Linode Instance '%s': %v", id, err)
		}
	}
}
