# Introduction

## Module Title: Production Grade Deployments - Deep Dive

## Module info

* **Duration:** ~5 hours 
* **Works with k8s version** 1.24
* **Workshop Domain:** *.workshop.metakube.org
* **Optimized for** Linux and MacOS
* Windows Users use WSL and/or PowerShell (you must be able to use shell variables)

## Prerequisites for trainers

* rotate AWS route53 token for external-dns and cert-manager
* AWS Route53: remove old DNS entries (*.workshop.metakube.org)
* Cluster setup via Terraform
* [Create 3 local storage worker nodes](../850-mysql-local-storage/README.md#prerequisites) for mysql example

## Prerequisites for participants

You need to have these tools installed on your local machine:

- kubectl (>1.24)
- helm (>3.10)
- helmfile
- Optional: kubectx

## Preparation for participants

* Before you begin with the actual exercise please make sure to follow these steps to work in your own environment:

  ```shell
  read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
  export YOURNAME
  kubectl create ns ${YOURNAME}
  kubectl label namespace ${YOURNAME} deepdive-pgd=true
  kubectl config set-context --current --namespace=${YOURNAME}
  ```

* Clone this repository to your working station and change into the directory for the following exercises

## Participants start with

Exercise [100-deployment-strategy](../100-deployment-strategy/README.md)

## Workshop environment

You will share a single Kubernetes cluster with other workshop participants.
Your own examples can be executed and tested in you personal namespace.
