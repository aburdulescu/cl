#!/bin/bash

#set -x

create_asset() {
    if [[ $# -ne 3 ]]
    then
        echo "error: create_asset() needs three arguments: VERSION, OS and CPUARCH "
        exit 1
    fi

    VERSION=$1
    OS=$2
    CPUARCH=$3

    ASSETNAME=$VERSION"_"$OS-$CPUARCH
    ASSETDIR=assets/$VERSION
    FILENAME=$ASSETDIR/"cl_"$ASSETNAME".tar.gz"

    echo "Creating release asset: "$ASSETNAME

    # create assets dir
    mkdir -p $ASSETDIR

    # build stripped binary
    GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w -X main.version="$ASSETNAME

    if [[ $OS == "windows" ]]
    then
        BINARY_NAME=cl.exe
    else
        BINARY_NAME=cl
    fi
    # compress binary
    tar -czvf $FILENAME $BINARY_NAME

    # calculate hash of binary
    sha512sum $FILENAME > $FILENAME.sha512
}

main() {
    if [[ $# -ne 1 ]]
    then
        echo "error: main() needs one argument: VERSION"
        exit 1
    fi

    VERSION=$1
    OSES="linux darwin windows"
    ARCHES="386 amd64"

    # clean dir before building
    rm -rf assets/$VERSION

    for os in $OSES
    do
        for arch in $ARCHES
        do
            create_asset $VERSION $os $arch
        done
    done
}

main $@
