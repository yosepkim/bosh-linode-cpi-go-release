module bosh-linode-cpi

go 1.13

require (
	github.com/bmatcuk/doublestar v1.2.2 // indirect
	github.com/charlievieth/fs v0.0.0-20170613215519-7dc373669fa1 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.0-20200104100158-8c5bf9093331
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/linode/linodego v0.12.2
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.8.1
	github.com/stretchr/testify v1.4.0
	github.com/vektra/mockery v0.0.0-20181123154057-e78b021dcbb5 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace bosh-linode-cpi => ../
