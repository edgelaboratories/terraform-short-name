JUNIT_REPORT ?= junit.xml
BUILD_TAG ?= terraform-modules-$(shell id -un)
AWS_KEY_NAME ?= jenkins-builds
TEST_ARGS ?= -v ./...

export TF_VAR_stack_id ?= $(BUILD_TAG)
export TF_VAR_aws_key_name ?= $(AWS_KEY_NAME)

export USERNAME ?= admin
export TF_VAR_username = $(USERNAME)

all: test

test:
	go test -count 1 -timeout 0 ${TEST_ARGS}
