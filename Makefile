.PHONY: all test clean images
.DEFAULT_GOAL := all

# Boiler plate for bulding Docker containers.
# All this must go at top of file I'm afraid.
IMAGE_PREFIX := quay.io/weaveworks/cortex-
IMAGE_TAG := $(shell ./tools/image-tag)
UPTODATE := .uptodate

# Building Docker images is now automated. The convention is every directory
# with a Dockerfile in it builds an image calls quay.io/weaveworks/<dirname>.
# Dependencies (i.e. things that go in the image) still need to be explicitly
# declared.
%/$(UPTODATE): %/Dockerfile
	$(SUDO) docker build -t $(IMAGE_PREFIX)$(shell basename $(@D)) $(@D)/
	$(SUDO) docker tag $(IMAGE_PREFIX)$(shell basename $(@D)) $(IMAGE_PREFIX)$(shell basename $(@D)):$(IMAGE_TAG)
	touch $@

# Get a list of directories containing Dockerfiles
DOCKERFILES := $(shell find . -name tools -prune -o -name vendor -prune -o -type f -name 'Dockerfile' -print)
UPTODATE_FILES := $(patsubst %/Dockerfile,%/$(UPTODATE),$(DOCKERFILES))
DOCKER_IMAGE_DIRS := $(patsubst %/Dockerfile,%,$(DOCKERFILES))
IMAGE_NAMES := $(foreach dir,$(DOCKER_IMAGE_DIRS),$(patsubst %,$(IMAGE_PREFIX)%,$(shell basename $(dir))))
images:
	$(info $(IMAGE_NAMES))
	@echo > /dev/null

# Generating proto code is automated.
PROTO_DEFS := $(shell find . -name tools -prune -o -name vendor -prune -o -type f -name '*.proto' -print)
PROTO_GOS := $(patsubst %.proto,%.pb.go,$(PROTO_DEFS))

# Building binaries is now automated.  The convention is to build a binary
# for every directory with main.go in it, in the ./cmd directory.
MAIN_GO := $(shell find . -name tools -prune -o -name vendor -prune -o -type f -name 'main.go' -print)
EXES := $(foreach exe, $(patsubst ./cmd/%/main.go, %, $(MAIN_GO)), ./cmd/$(exe)/$(exe))
GO_FILES := $(shell find . -name tools -prune -o -name vendor -prune -o -name cmd -prune -o -type f -name '*.go' -print)
define dep_exe
$(1): $(dir $(1))/main.go $(GO_FILES) $(PROTO_GOS)
$(dir $(1))$(UPTODATE): $(1)
endef
$(foreach exe, $(EXES), $(eval $(call dep_exe, $(exe))))

# Manually declared dependancies And what goes into each exe
cortex.pb.go: cortex.proto
ring/ring.pb.go: ring/ring.proto
all: $(UPTODATE_FILES)
test: $(PROTO_GOS)

# And now what goes into each image
build/$(UPTODATE): build/*

# All the boiler plate for building golang follows:
SUDO := $(shell docker info >/dev/null 2>&1 || echo "sudo -E")
BUILD_IN_CONTAINER := true
RM := --rm
GO_FLAGS := -ldflags "-extldflags \"-static\" -linkmode=external -s -w" -tags netgo -i
NETGO_CHECK = @strings $@ | grep cgo_stub\\\.go >/dev/null || { \
       rm $@; \
       echo "\nYour go standard library was built without the 'netgo' build tag."; \
       echo "To fix that, run"; \
       echo "    sudo go clean -i net"; \
       echo "    sudo go install -tags netgo std"; \
       false; \
}

ifeq ($(BUILD_IN_CONTAINER),true)

$(EXES) $(PROTO_GOS) lint test shell: build/$(UPTODATE)
	@mkdir -p $(shell pwd)/.pkg
	$(SUDO) time docker run $(RM) -ti \
		-v $(shell pwd)/.pkg:/go/pkg \
		-v $(shell pwd):/go/src/github.com/weaveworks/cortex \
		$(IMAGE_PREFIX)build $@

configs-integration-test: build/$(UPTODATE)
	@mkdir -p $(shell pwd)/.pkg
	DB_CONTAINER="$$(docker run -d -e 'POSTGRES_DB=configs_test' postgres:9.6)"; \
	$(SUDO) docker run $(RM) -ti \
		-v $(shell pwd)/.pkg:/go/pkg \
		-v $(shell pwd):/go/src/github.com/weaveworks/cortex \
		-v $(shell pwd)/cmd/configs/migrations:/migrations \
		--workdir /go/src/github.com/weaveworks/cortex \
		--link "$$DB_CONTAINER":configs-db.cortex.local \
		$(IMAGE_PREFIX)build $@; \
	status=$$?; \
	test -n "$(CIRCLECI)" || docker rm -f "$$DB_CONTAINER"; \
	exit $$status

else

$(EXES): build/$(UPTODATE)
	go build $(GO_FLAGS) -o $@ ./$(@D)
	$(NETGO_CHECK)

%.pb.go: build/$(UPTODATE)
	protoc -I ./vendor:./$(@D) --gogoslick_out=plugins=grpc:./$(@D) ./$(patsubst %.pb.go,%.proto,$@)

lint: build/$(UPTODATE)
	./tools/lint -notestpackage -ignorespelling queriers -ignorespelling Queriers .

test: build/$(UPTODATE)
	./tools/test -netgo

shell: build/$(UPTODATE)
	bash

configs-integration-test:
	/bin/bash -c "go test -tags netgo,integration -timeout 30s ./configs/..."

endif

clean:
	$(SUDO) docker rmi $(IMAGE_NAMES) >/dev/null 2>&1 || true
	rm -rf $(UPTODATE_FILES) $(EXES) $(PROTO_GOS)
	go clean ./...

clean-vendor:
	find vendor -type f ! -name '*.go' -a ! -name '*.s' -a ! -name '*.h' -a ! -name '*.proto' -a ! -name LICENSE | xargs -L 10 rm
	find vendor -type f -name '*_test.go' | xargs -L 10 rm
