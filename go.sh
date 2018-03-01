#!/bin/bash

CHECKIN=${1:-""}
BRANCH=${2:-"master"}
FILES=${3:-"*.*"}

echo "formatting and building $FILES"
gofmt -w ./$FILES && go build ./$FILES 

if [ "$CHECKIN" != "" ]; then 
 echo "Checking in $FILES into $BRANCH"
 git add $FILES
 git commit -m "\"${CHECKIN}\""
 go get -u github.com/efidoman/xdripgo
fi



