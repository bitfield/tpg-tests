#!/bin/bash

for TOPIC in */; do
  for LISTING in ${TOPIC}[0-9+]; do
    SCRIPT=${LISTING}-test.txtar
    cp ${LISTING}.txtar $SCRIPT || exit
    txtar-c $LISTING >>$SCRIPT
    echo -n "${LISTING}... " && testscript $SCRIPT
    rm $SCRIPT
  done
done