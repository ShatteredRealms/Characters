#####################################################################################
#   _____ _           _   _                    _   _____            _               #
#  / ____| |         | | | |                  | | |  __ \          | |              #
# | (___ | |__   __ _| |_| |_ ___ _ __ ___  __| | | |__) |___  __ _| |_ __ ___  ___ #
#  \___ \| '_ \ / _` | __| __/ _ \ '__/ _ \/ _` | |  _  // _ \/ _` | | '_ ` _ \/ __|#
#  ____) | | | | (_| | |_| ||  __/ | |  __/ (_| | | | \ \  __/ (_| | | | | | | \__ \#
# |_____/|_| |_|\__,_|\__|\__\___|_|  \___|\__,_| |_|  \_\___|\__,_|_|_| |_| |_|___/#
#                   _____ _                          _                              #
#                  / ____| |                        | |                             #
#                 | |    | |__   __ _ _ __ __ _  ___| |_ ___ _ __ ___               #
#                 | |    | '_ \ / _` | '__/ _` |/ __| __/ _ \ '__/ __|              #
#                 | |____| | | | (_| | | | (_| | (__| ||  __/ |  \__ \              #
#                  \_____|_| |_|\__,_|_|  \__,_|\___|\__\___|_|  |___/              #
#####################################################################################

#
# Makefile for building, running, and testing
#

# Import dotenv
ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif

# Application versions
BASE_VERSION = 0.0.1

# Gets the directory containing the Makefile
BUILD_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Root code directory
ROOT_DIR = $(realpath $(BUILD_DIR)/..)

# Directory containing applications
BASE_APP_DIR = $(realpath $(ROOT_DIR)/cmd)

# Binary output directory
BIN_DIR = $(realpath $(ROOT_DIR)/bin)

# Entrypoint for applications
APP_MAIN = $(BASE_APP_DIR)/characters/main.go

# Base container registry
SRO_BASE_REGISTRY ?= 779965382548.dkr.ecr.us-east-1.amazonaws.com
SRO_REGISTRY ?= $(SRO_BASE_REGISTRY)/sro
BASE_TAG = sro-chracters

# The registry for this service
REGISTRY = $(SRO_REGISTRY)/characters
time=$(shell date +%s)

#   _____                    _
#  |_   _|                  | |
#    | | __ _ _ __ __ _  ___| |_ ___
#    | |/ _` | '__/ _` |/ _ \ __/ __|
#    | | (_| | | | (_| |  __/ |_\__ \
#    \_/\__,_|_|  \__, |\___|\__|___/
#                  __/ |
#                 |___/

build:
	go build -o $(BIN_DIR)/characters $(APP_MAIN)

test:
	ginkgo $(ROOT_DIR)/... -covermode atomic

report: test
	go tool cover -func=$(ROOT_DIR)/coverage.out
	go tool cover -html=$(ROOT_DIR)/coverage.out

run-local:
	SRO_DB_FILE=$(ROOT_DIR)/test/db.yaml SRO_KEY_DIR=$(ROOT_DIR)/test/auth SRO_CHARACTERS_PORT=8081 go run $(APP_MAIN)

deploy: aws-docker-login push

build-image:
	docker build -t $(BASE_TAG) -f Dockerfile ..


aws-docker-login:
	aws ecr get-login-password | docker login --username AWS --password-stdin $(SRO_BASE_REGISTRY)

push-prod: build-image
	docker tag $(BASE_TAG) $(REGISTRY):latest
	docker tag $(BASE_TAG) $(REGISTRY):$(BASE_VERSION)
	docker tag $(BASE_TAG) $(REGISTRY):$(BASE_VERSION)-$(time)
	docker push $(REGISTRY):latest
	docker push $(REGISTRY):$(BASE_VERSION)
	docker push $(REGISTRY):$(BASE_VERSION)-$(time)

push-%: build-image
	docker tag $(BASE_TAG) $(REGISTRY)/$*:latest
	docker tag $(BASE_TAG) $(REGISTRY)/$*:$(BASE_VERSION)
	docker tag $(BASE_TAG) $(REGISTRY)/$*:$(BASE_VERSION)-$(time)
	docker push $(REGISTRY)/$*:latest
	docker push $(REGISTRY)/$*:$(BASE_VERSION)
	docker push $(REGISTRY)/$*:$(BASE_VERSION)-$(time)
