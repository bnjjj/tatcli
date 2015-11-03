#!/bin/bash

for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64 arm; do
        if [[ $GOARCH == "arm" && $GOOS != "linux" ]]; then
          continue;
        fi;
        architecture="${GOOS}-${GOARCH}"
        echo "Building ${architecture} ${path} update url : ${URL_UPDATE_RELEASE}"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go build -ldflags "-X ${PROJECT_PATH}/${PROJECT_NAME}/update.architecture=${architecture} -X ${PROJECT_PATH}/${PROJECT_NAME}/update.urlUpdateRelease=${URL_UPDATE_RELEASE} -X ${PROJECT_PATH}/${PROJECT_NAME}/update.urlUpdateSnapshot=${URL_UPDATE_SNAPSHOT}" -o bin/tatcli-${architecture}
    done
done
