#!/bin/bash

for TOPIC in */; do
  for LISTING in ${TOPIC}[0-9+]; do
    # cd $LISTING && go get -t -u && go mod tidy && cd -
    ./run_testscript.sh $LISTING
  done
done