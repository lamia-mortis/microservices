#!/bin/sh

set -x

WORK_DIR=${WORK_DIR}
TMP_DIR=${TMP_DIR}
GATEWAY_BINARY_FILE=${GATEWAY_BINARY_FILE}

if [ ! -d $WORK_DIR ]; then
    mkdir -p $WORK_DIR
fi

if [ ! -z $WORK_DIR ]; then
    cp -r $TMP_DIR/* $WORK_DIR
fi 

$WORK_DIR/$GATEWAY_BINARY_FILE