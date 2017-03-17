#!/bin/bash

set -e

script_root=`dirname "${BASH_SOURCE}"`
source $script_root/env.sh


echo "create Mole Service..."
cat mole-service-template.yaml | $KUBECTL $KUBECTL_OPTIONS create -f -


echo "create Mole replicationcontroller..."
cat mole-controller-template.yaml | $KUBECTL $KUBECTL_OPTIONS create -f -

