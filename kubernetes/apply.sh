#!/bin/sh

NAMESACE=bellinghamcodes

kubectl --namespace=${NAMESACE} apply -f ./secrets.yaml
kubectl --namespace=${NAMESACE} apply -f ./deployments.yaml
