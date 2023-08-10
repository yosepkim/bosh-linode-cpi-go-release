package instance

type Service interface {
	Create(vmProps *Properties, registryEndpoint string) (string, error)
	Delete(id string) error
	CleanUp(id string)
}

type Properties struct {
	LinodeType string
	Region     string
	Tags       Tags
}

type Tags []string
