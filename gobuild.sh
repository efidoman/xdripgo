#!/bin/bash

FILES=${1:-"*.go"}
CHECKIN=${2:-""}
BRANCH=${3:-"master"}

echo "formatting and building $FILES"
gofmt -w ./$FILES && go build ./$FILES 

if [ "$CHECKIN" != "" ]; then 
 echo "Checking in $FILES into $BRANCH"
 git add $FILES
 git commit -m "\"${CHECKIN}\""
fi

./go.sh


