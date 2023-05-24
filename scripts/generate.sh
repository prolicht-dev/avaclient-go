#!/bin/bash

####################################################################
# Config section 
####################################################################

API_URL=https://avacloud-api.dangl-it.com/swagger/swagger.json
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
LIB_DIR=$SCRIPT_DIR/../
cd $SCRIPT_DIR

####################################################################
# Script section 
####################################################################

echo "[I] Checking availability of JAVA"
if type -p java >/dev/null 2>&1; then
    java_version=$(java -version 2>&1 | awk -F '"' '/version/ {print $2}')
    echo "[I] found java executable in PATH: $java_version" 
elif [[ -n "$JAVA_HOME" ]] && [[ -x "$JAVA_HOME/bin/java" ]];  then
    java_version=$("$_java" -version 2>&1 | awk -F '"' '/version/ {print $2}')
    echo "[I] found java executable in JAVA_HOME: $java_version" 
else
    echo "[E] no java executable found" 
    exit 1
fi

generater_jar=openapi-generator-cli-6.6.0.jar
generator_version=$(java -jar "$SCRIPT_DIR"/$generater_jar version)
echo "[I] Running openapi generator $generator_version for URL: $API_URL"

api_version=$(curl $API_URL 2>&1 | grep -o -E "\"version\":\s*\"[v0-9 \\.-]+\"" | awk -F\: '{print $2}' | awk -F'"' '{print $2}')
api_main_version=$(echo $api_version | awk -F' ' '{print $1}')
echo "[I] API version: $api_version, main version: $api_main_version"

java -jar "$SCRIPT_DIR"/$generater_jar generate -i $API_URL -g go \
    --additional-properties=packageName=avaclient,packageVersion=$api_main_version \
    --git-host github.com --git-user-id prolicht-dev --git-repo-id avaclient-go \
    --api-package api --model-package model \
    -p enumClassPrefix=true \
    -o $LIB_DIR

echo "[I] Applying fix for polymorphic IElement"
sed -i 's/\[\]IElementDto/\[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_dto.go
sed -i 's/\[\]IElementDto/\[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_group_dto.go
sed -i 's/\[\]IElementDto/\[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_group_dto_all_of.go

cd $LIB_DIR

echo "[I] Updating all dependencies"
go get -u ./...

echo "[I] Cleaning up go.mod"
go mod tidy

echo "[I] Running go fmt: $(go version)"
go fmt ./...

echo "[I] DONE"
