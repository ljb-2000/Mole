#!/bin/bash

script_root=`dirname "${BASH_SOURCE}"`
source $script_root/env.sh


echo "Stopping Mole replicationcontroller..."
$KUBECTL $KUBECTL_OPTIONS delete replicationcontroller mole

echo "Deleting Mole service..."
$KUBECTL $KUBECTL_OPTIONS delete service mole

