package action

import (
	"bosh-linode-cpi/config"

	"bosh-linode-cpi/linode/client"
	boshlinodeconfig "bosh-linode-cpi/linode/config"
	instance "bosh-linode-cpi/linode/instance_service"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/linode/linodego"
)

var LinodeClientFunc func(boshlinodeconfig.LinodeConfig, boshlog.Logger) (linodego.Client, error) = client.NewLinodeClient

type ConcreteFactory struct {
	cfg    config.Config
	logger boshlog.Logger
}

func NewConcreteFactory(
	cfg config.Config,
	logger boshlog.Logger,
) ConcreteFactory {
	return ConcreteFactory{
		cfg,
		logger}
}

func (f ConcreteFactory) Create(method string) (Action, error) {
	linodeClient, err := LinodeClientFunc(f.cfg.Cloud.Properties.Linode, f.logger)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Building linode client")
	}

	vmService := instance.NewLinodeInstanceService(
		linodeClient,
		f.logger,
	)

	actions := map[string]Action{
		"create_vm": NewCreateVM(vmService),
		"delete_vm": NewDeleteVM(vmService),
	}

	action, found := actions[method]
	if !found {
		return nil, bosherr.Errorf("Could not create action with method %s", method)
	}

	return action, nil
}
