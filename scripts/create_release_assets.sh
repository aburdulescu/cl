#!/bin/bash

if [[ $# -ne 3 ]]
then
    echo "error: need VERSION, OS and CPUARCH "
    exit 1
fi

VERSION=$1
OS=$2
CPUARCH=$3

ASSETNAME=$VERSION"_"$OS-$CPUARCH
ASSETDIR=assets/$VERSION/$OS/$CPUARCH
FILENAME=$ASSETDIR/"cl_"$ASSETNAME".tar.gz"

echo "Creating release asset: "$ASSETNAME

# clean asset dir and create it again
rm -rf $ASSETDIR
mkdir -p $ASSETDIR

# build stripped binary
go build -ldflags="-s -w" > /dev/null

# compress binary
tar -czvf $FILENAME cl

# calculate hash of binary
sha512sum $FILENAME > $FILENAME.sha512
