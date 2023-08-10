package instance

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/linode/linodego"
)

const linodeInstanceServiceLogTag = "LinodeInstanceService"

type LinodeInstanceService struct {
	linodeClient linodego.Client
	logger       boshlog.Logger
}

func NewLinodeInstanceService(
	linodeClient linodego.Client,
	logger boshlog.Logger,
) LinodeInstanceService {
	return LinodeInstanceService{
		linodeClient: linodeClient,
		logger:       logger,
	}
}
