#!/bin/bash

# Most clusters will just be accessed with 'kubectl' on $PATH.
# However, some might require a different command. For example, GKE required
# KUBECTL='gcloud container kubectl' for a while. Now that most of our
# use cases just need KUBECTL=kubectl, we'll make that the default.
KUBECTL=${KUBECTL:-kubectl}

# Kubernetes api address for $KUBECTL 
# The default value is 127.0.0.1:8080
# When the Kubernetes api server is not local, We can easily access the api by edit KUBERNETES_API_SERVER's value
KUBERNETES_API_SERVER=${KUBERNETES_API_SERVER:-'192.168.180.101:8080'}

# Kubernetes namespace for mole and components.
NAME_SPACE=${NAME_SPACE:-'default'}

# Kubernetes options config
KUBECTL_OPTIONS="--namespace=$NAME_SPACE --server=$KUBERNETES_API_SERVER"
