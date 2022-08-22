#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: $0 example_path"
    exit 1
fi

LISTING=$1
SCRIPT=${LISTING}-test.txtar
txtar-c -quote $LISTING >>$SCRIPT.tmp
# Ensure that any 'unquote' directives come before the actual script
grep unquote $SCRIPT.tmp >$SCRIPT
cat ${LISTING}/test.txtar >>$SCRIPT || exit
echo >>$SCRIPT
grep -v 'unquote' $SCRIPT.tmp >>$SCRIPT
rm $SCRIPT.tmp
echo -n "${LISTING}... " && testscript $SCRIPT || exit
rm $SCRIPT
