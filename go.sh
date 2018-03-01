#!/bin/bash

CHECKIN=${1:-""}
BRANCH=${2:-"master"}
SRC=${3:-"*.go"}
FILES=${3:-"*.*"}

echo "formatting and building $SRC"
gofmt -w ./$SRC && go build ./$SRC 

if [ "$CHECKIN" != "" ]; then 
 echo "Checking in $FILES into $BRANCH"
 git add $FILES
 git commit -m "\"${CHECKIN}\""
 go get -u github.com/efidoman/xdripgo
fi



