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

generator_version=$(java -jar "$SCRIPT_DIR"/openapi-generator-cli.jar version)
echo "[I] Running openapi generator $generator_version for URL: $API_URL"

api_version=$(curl $API_URL 2>&1 | grep -o -E "\"version\":\s*\"[v0-9 \\.-]+\"" | awk -F\: '{print $2}' | awk -F'"' '{print $2}')
api_main_version=$(echo $api_version | awk -F' ' '{print $1}')
echo "[I] API version: $api_version, main version: $api_main_version"

java -jar "$SCRIPT_DIR"/openapi-generator-cli.jar generate -i $API_URL -g go \
    --additional-properties=packageName=avaclient,packageVersion=$api_main_version \
    --git-host github.com --git-user-id prolicht-dev --git-repo-id avaclient-go \
    --api-package api --model-package model \
    -p enumClassPrefix=true \
    -o $LIB_DIR

echo "[I] Applying fix for polymorphic IElement"
sed -i '/type IElementDto struct {/,/}/ctype IElementDto interface {\n    isIElementDto()\n}' ${LIB_DIR}model_i_element_dto.go
sed -i 's/Elements \[\]IElementDto/Elements \[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_dto.go
sed -i 's/Elements \[\]IElementDto/Elements \[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_group_dto.go
sed -i 's/Elements \[\]IElementDto/Elements \[\]IElementDtoHolder/g' ${LIB_DIR}model_service_specification_group_dto_all_of.go

echo "[I] Applying fix for content-type check"
sed -i 's/vnd\\.\[\^;\]\+/(?:vnd\\.[^;]+|problem)/g' ${LIB_DIR}client.go

## GO FMT
GODIRS=$(go list $LIB_DIR/... | grep -v /vendor/)
GOFILES=$(find $LIB_DIR -type f -name '*.go' -not -path "./vendor/*")

echo "[I] Running goimports: $(go version)"
GOIMPORTS=goimports
if ! [ -x "$(command -v goimports)" ]; then
    go get golang.org/x/tools/cmd/goimports
    GOIMPORTS=~/go/bin/goimports
fi
$GOIMPORTS -w $GOFILES

echo "[I] Running go fmt: $(go version)"
go fmt $GODIRS

echo "[I] Running go tests: $(go version)"
go test $GODIRS

echo "[I] Cleaning up..."
rm -f $LIB_DIR/git_push.sh
rm -rf $LIB_DIR/.openapi-generator
rm -rf $LIB_DIR/api

echo "[I] DONE"
