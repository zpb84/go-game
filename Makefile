GOOS	=linux
GOARCH=amd64

PROJECT_PATH=$(shell pwd)
DATE=$(shell date +"%Y/%m/%d")
BUILD_PATH=$(PROJECT_PATH)/build/app

MAJOR_VERSION=$(shell git tag | head -1)
COMMIT_NUMBER=$(shell git rev-list --count HEAD)
COMMIT_SHA=$(shell git rev-parse --short=8 HEAD)
VERSION=v.0.$(MAJOR_VERSION).$(COMMIT_NUMBER)

build: clean
	go build --ldflags "-X 'main.version=$(VERSION)' -X 'main.date=$(DATE)' -X 'main.commit=$(COMMIT_SHA)'" -o $(BUILD_PATH) $(PROJECT_PATH)/cmd

clean: show
	rm -f $(BUILD_PATH)

show:
	@echo "******** BUILD ENVIRONMENT ********"
	@echo "VERSION:                 $(VERSION)"
	@echo "GOPATH:                  $(GOPATH)"
	@echo "GOOS:                    $(GOOS)"
	@echo "GOARCH:                  $(GOARCH)"
	@echo ""

run: build
	$(BUILD_PATH)