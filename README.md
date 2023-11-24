# K8s Deep Dive 

## Mastering Production-Grade Deployments with Practical Insights and Hands-On Exercises

Workshop by [SysEleven](https://syseleven.de)

![SysEleven](https://www.syseleven.de/wp-content/uploads/2020/10/SysEleven_XL_Logo_quer_RGB.png)

---

### Event Website

[ContainerDay Workshops 2023](https://www.containerdays.io/containerday-workshops-2023/#k8s-deep-dive-mastering-production-grade-deployments-with-practical-insights-and-hands-on-exercises) in Munich

### Description

In this workshop you will get a deeper insight into production grade deployments. A mixture of theory, practical examples and exercises including a test environment is waiting for you. During the workshop, participants will be provided with ready-made Kubernetes clusters. This course is aimed at DevOps, Cloud Platform Operations, Cloud Admins. Participants should have the following prior knowledge:
- Kubernetes basics
- You have already had hands-on experience running on Kubernetes
- Fundamentals of Helm charts (gotemplates, creating and running container images)
- You are basically capable of rolling out helm charts in Kubernetes
- Interest in cloud-native technologies

### Requirements

- your private notebook (MacOS / Linux recommended)
- shell console
- kubectl client installed (v1.25+) - [Installation instructions](https://kubernetes.io/docs/tasks/tools/#kubectl)
- helm client installed (v3.10+) - [Installation instructions](https://helm.sh/docs/intro/install/)

### Further Course Instructions

Will be updated prior to the workshop date...

### Docker Images used

```text
nginx:1.24.0
nginx:1.25.0
gcr.io/google_containers/hpa-example
busybox
nginxinc/nginx-unprivileged:1.25-alpine
nginx
syseleven/metakube-hello:1.0.0
docker.io/bitnami/external-dns:0.13.4-debian-11-r11
registry.k8s.io/ingress-nginx/controller:v1.7.0@sha256:7612338342a1e7b8090bef78f2a04fffcadd548ccaabe8a47bf7758ff549a5f7
docker.io/hashicorp/http-echo
quay.io/jetstack/cert-manager-controller:v1.11.1
quay.io/jetstack/cert-manager-cainjector:v1.11.1
quay.io/jetstack/cert-manager-webhook:v1.11.1
```
