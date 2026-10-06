#!/bin/bash

# exit immediately if any command fails
set -e

# verify kubectl and k3d are installed
kubectl version --client
echo "kubectl is installed"

k3d --version
echo "k3d is installed"

# create a k3d cluster with an explicit name and 2 agents
k3d cluster create mycluster -a 2

# show cluster info and running containers
docker ps
kubectl cluster-info

# create the deployment from DockerHub
kubectl create deployment todo-app --image=yehezkel1086/todo-app

# wait for the deployment to roll out successfully
kubectl rollout status deployment/todo-app

# display current pods and deployments
kubectl get pods
kubectl get deployments

# 7. stream logs dynamically using the label selector
kubectl logs -f -l app=todo-app
