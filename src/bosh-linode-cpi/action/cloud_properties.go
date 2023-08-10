package action

type Environment map[string]interface{}

type VMCloudProperties struct {
	LinodeType string   `json:"linode_type,omitempty"`
	Region     string   `json:"region,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

func (n VMCloudProperties) Validate() error {
	return nil
}
