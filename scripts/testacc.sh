#!/bin/bash


Title() {
    printf -v Bar '%*s' $((${#1} + 4)) ' ' 
    printf '%s\n| %s |\n%s\n' "${Bar// /-}" "$1" "${Bar// /-}"
}

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

ENV="/opt/atricore/tools/env.sh"
if [ ! -f $ENV ] ; then
    echo "File not found : $ENV"
    exit 1
fi

source $ENV

# This is the pipeline folder name for the JOSSO EE server (i.e. josso-ee-2-m)
if [ -z "$1" ] || [ -z "$2" ] ; then
    echo "Usage: $0 <pipeline> <josso-version>"
    exit 1
fi

PIPELINE=$1
JOSSO_VERSION=$2

BUILD_FOLDER=$SCRIPT_DIR/../../$PIPELINE
if [ ! -d "$BUILD_FOLDER" ] ; then
        echo "Build folder not found : $BUILD_FOLDER"
        exit 1
fi
if [ ! -f "$BUILD_FOLDER/buildNumber.properties" ] ; then
        echo "Build number not found : $BUILD_FOLDER/buildNumber.properties"
        exit 1
fi

BUILD_DATE=`sed 's/^#\(.*\)/\1/;2q;d' "$BUILD_FOLDER"/buildNumber.properties`
BUILD_NUM=`sed 's/^buildNumber=\(.*\)/\1/;3q;d' "$BUILD_FOLDER"/buildNumber.properties`
BUILD_TIMESTAMP=$(date -d "$BUILD_DATE" +%Y%m%d.%k%M%S)
BUILD_TAG=$PIPELINE-$BUILD_TIMESTAMP-$BUILD_NUM
MVN_TARGET="$BUILD_FOLDER"/distributions/josso-ee/target
TMP=/tmp/josso

export JOSSO_API_SECRET="7oUHlv(HLT%vxK4L"
export JOSSO_API_CLIENT_ID="idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
export JOSSO_API_ENDPOINT="http://localhost:8111/atricore-rest/services"
export JOSSO_API_USERNAME="admin"
export JOSSO_API_PASSWORD="atricore"

LOG_FILE=/tmp/josso-sdk-acctest-$BUILD_TAG.log
# iamtf_acctest_server      2.5.3-1
DOCKER_IMG="atricore/josso:$JOSSO_VERSION-$BUILD_NUM"

# Start docker container

cd $SCRIPT_DIR/../docker/
if [ ! -d $SCRIPT_DIR/../docker ] ; then 
    echo "Invalid docker folder"
    exit 1
fi

DOCKER_IMAGE="atricore/josso:$JOSSO_VERSION-$BUILD_NUM"
DOCKER_CONTAINER="terraform-provider-josso-server-$JOSSO_VERSION-$BUILD_NUM"

#Title "Building docker image : $DOCKER_IMAGE"
#docker build --build-arg "IAMGE=$DOCKER_IMAGE"

Title "Building docker container : $DOCKER_CONTAINER" 
docker run \
        --name "$DOCKER_CONTAINER" \
        --detach \
        --env JOSSO_CLIENT_ID="$JOSSO_API_CLIENTID" \
        --env JOSSO_CLIENT_SECRET="$JOSSO_API_SECRET" \
        --env JOSSO_ADMIN_USR=myadmin \
        --env JOSSO_ADMIN_PWD=changeme \
        --env JOSSO_SKIP_ADMIN_CREATE=false \
        --env KARAF_DEBUG=true \
        -p8111:8081 -p8222:8101 -p8444:8443 \
        "$DOCKER_IMAGE"


# TODO : Wait for the server to start!
for i in {1..15}; do
    curl -I "$JOSSO_API_ENDPOINT/info" | grep "HTTP/1.1 401 Unauthorized"
    if [ $? -eq 0 ] ; then
        break
    fi
    sleep 5
done

cd $SCRIPT_DIR/..

Title "Running tests : $SCRIPT_DIR/.."
make testacc
MAKE_STATUS=$?


Title "Destroying docker container : $DOCKER_CONTAINER" 
docker rm -f "$DOCKER_CONTAINER"

if [ $MAKE_STATUS -ne 0 ] ; then
    Title "There are TESTS errors (make status): $MAKE_STATUS"
    exit 1
fi
