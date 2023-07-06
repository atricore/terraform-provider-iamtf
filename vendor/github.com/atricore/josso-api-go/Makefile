GENERATOR=/opt/atricore/tools/openapi-generator-cli/openapi-generator-cli.sh
OPENAPI_GENERATOR_VERSION=6.2.1
#GENERATOR=/data/atricore/tools/openapi-generator-cli/openapi-generator-cli.sh
#SWAGGER_FILE=~/.m2/repository/com/atricore/idbus/console/console-api/1.4.3-SNAPSHOT/console-api-1.4.3-SNAPSHOT-swagger.yaml
#SWAGGER_FILE=./console-api-1.4.3-SNAPSHOT-swagger.yaml

SWAGGER_FILE=./console-api-1.5.1-SNAPSHOT-swagger.json

PGK_NAME=jossoappi

default: all

build:
	go install

dep: # Download required dependencies
	go mod tidy
	go mod vendor

test:
	go test

generate:
	OPENAPI_GENERATOR_VERSION=${OPENAPI_GENERATOR_VERSION} $(GENERATOR) generate -i $(SWAGGER_FILE) -g go -o . --additional-properties=packageName=$(PGK_NAME) --additional-properties=disallowAdditionalPropertiesIfNotPresent=false --git-repo-id=josso-api-go --git-user-id=atricore

all: generate dep build
