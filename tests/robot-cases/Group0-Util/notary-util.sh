#!/bin/bash

CMD=$1
HOST=$2
PROJECT=$3
IMAGE=$4
TAG=$5
NOTARY_SERVER_ENDPOINT=$6
USER=$7
PWD=$8
PASSHRASE=$8

export DOCKER_CONTENT_TRUST=1

export NOTARY_ROOT_PASSPHRASE=$PASSHRASE
export NOTARY_TARGETS_PASSPHRASE=$PASSHRASE
export NOTARY_SNAPSHOT_PASSPHRASE=$PASSHRASE
export DOCKER_CONTENT_TRUST_ROOT_PASSPHRASE=$PASSHRASE
export DOCKER_CONTENT_TRUST_REPOSITORY_PASSPHRASE=$PASSHRASE
export DOCKER_CONTENT_TRUST_OFFLINE_PASSPHRASE=$PASSHRASE
export DOCKER_CONTENT_TRUST_TAGGING_PASSPHRASE=$PASSHRASE

export NOTARY_AUTH=$(echo $USER:$PWD | base64)
echo $USER:$PWD

NOTARY_CMD_OPTIONS="notary -s https://$NOTARY_SERVER_ENDPOINT -d /root/.docker/trust"
if [ "$CMD" == "key_rotate" ]; then
    echo "$NOTARY_CMD_OPTIONS key rotate  $HOST/$PROJECT/$IMAGE snapshot -r"
    $NOTARY_CMD_OPTIONS key rotate  $HOST/$PROJECT/$IMAGE snapshot -r
elif  [ "$CMD" == "remove" ]; then
    echo "$NOTARY_CMD_OPTIONS remove -p $HOST/$PROJECT/$IMAGE $TAG"
    $NOTARY_CMD_OPTIONS remove -p $HOST/$PROJECT/$IMAGE $TAG
fi
