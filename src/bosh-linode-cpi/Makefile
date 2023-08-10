default: test

$(GOPATH)/bin/ginkgo:
	go get -u github.com/onsi/ginkgo/ginkgo

vet:
	go vet ./...

build:
	go build -o out/cpi bosh-linode-cpi/main

test: $(GOPATH)/bin/ginkgo vet build
	ginkgo -r -race -skipPackage=integration .

testint: check-token $(GOPATH)/bin/ginkgo vet
	LINODE_TOKEN=$(LINODE_TOKEN) \
	ginkgo integration -slowSpecThreshold=500 -progress -nodes=3 -randomizeAllSpecs -randomizeSuites $(GINKGO_ARGS) -v

# Checks and creates, if necessary, resources required to run integration tests.
configint: check-token
	cd ../../ci/test_infrastructure; terraform init; terraform apply -auto-approve \
	  -var 'linode_token="$(LINODE_TOKEN)"' \
	  -var 'prefix=cpi-integration'

# Deletes the resources created by the configint target
cleanint: check-token
	cd ../../ci/test_infrastructure; terraform init; terraform destroy -force \
	  -var 'linode_token="$(LINODE_TOKEN)"' \
	  -var 'prefix=cpi-integration'

	@echo "All integration tests resources have been removed. You will need to run the configint target before doing integration tests again."

check-token:
	@if test "$(LINODE_TOKEN)" = "" ; then \
	  echo "LINODE_TOKEN must be set"; \
	  exit 1; \
	fi

mocks:
	cd linode/instance_service; mockery -name=Service
