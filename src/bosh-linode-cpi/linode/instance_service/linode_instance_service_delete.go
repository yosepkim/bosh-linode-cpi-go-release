package instance

import (
	"context"
	"strconv"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (i LinodeInstanceService) Delete(id string) error {
	linodeId, err := strconv.Atoi(id)
	if err != nil {
		return bosherr.WrapErrorf(err, "Failed to find Linode Instance '%s' for delete", id)
	}
	i.logger.Debug(linodeInstanceServiceLogTag, "Deleting Linode Instance '%s'", id)
	err = i.linodeClient.DeleteInstance(context.Background(), linodeId)
	if err != nil {
		return bosherr.WrapErrorf(err, "Failed to delete Linode Instance '%s'", id)
	}

	return nil
}
