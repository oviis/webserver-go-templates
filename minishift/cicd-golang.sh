#!/bin/bash
readonly BASEDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd -P)"
readonly PROGRAM="$(basename "${BASH_SOURCE[0]}")"

# Fail on any errors
set -o errexit
# Enforce variables settings
set -o nounset

cd "${BASEDIR}"

##login as developer
oc login -u developer

##new project
oc new-project cicd-golang --description="will build and deploy an Golang webserver with templates" --display-name="cicd-golang"

#creating the project
oc project cicd-golang

##create an ephemeral jenkins master, if you want a persistent, please follow the docs
oc new-app jenkins-ephemeral

##waiting for jenkins master to be available, this cayn take some time
echo "....be passion, will wait up to 6min...., we need to install a jenkins-master here...."
sleep 350

##see if all resources are deployed and available
oc get all

