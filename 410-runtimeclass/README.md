# gvisor as RuntimeClass

## Introduction

Kubernetes can make use of a runtimeClass to run particular Pods on them.

The runtime gvisor (runsc) is a handler which works as a kernel proxy for the Pod.
Hereby it adds a layer of security around a Pod and filters syscalls made by the Pod.
Think of gvisor as of "a kernel in user space".

## Task

Install and use gvisor as a runtimeClass in your cluster.

**This installation is required only once per cluster. (Only 1 participant!)**

## Preparation

* Before you begin with the actual exercise please make sure to follow these steps to work in your own environment:

  ```shell
  # enter your name
  # example:
  # export YOURNAME=janedoe
  export YOURNAME=<YOURNAME>
  kubectl create ns ${YOURNAME}
  kubectl config set-context --current --namespace=${YOURNAME}
  ```

* Clone this repository to your working station and change into the directory for the following exercises

---

## Prerequisites

* a Kubernetes worker Node with direct SSH access or
* [node-shell plugin](https://github.com/kvaps/kubectl-node-shell) for kubectl

## Prepare a Node

### Taint and label

```shell
kubectl get nodes
kubcetl taint <NODENAME> runtime=gvisor:NoSchedule
kubectl label <NODENAME> runtime=gvisor
```

## Install gvisor on a dedicated Node

* Connect to the node via node-shell

```shell
kubectl node-shell <NODENAME>
```

* Get the latest release

```shell
(
  set -e
  ARCH=$(uname -m)
  URL=https://storage.googleapis.com/gvisor/releases/release/latest/${ARCH}
  wget ${URL}/runsc ${URL}/runsc.sha512 \
    ${URL}/containerd-shim-runsc-v1 ${URL}/containerd-shim-runsc-v1.sha512
  sha512sum -c runsc.sha512 \
    -c containerd-shim-runsc-v1.sha512
  rm -f *.sha512
  chmod a+rx runsc containerd-shim-runsc-v1
  sudo mv runsc containerd-shim-runsc-v1 /usr/local/bin
)
```

* Install runsc runtime

```shell
/usr/local/bin/runsc install
```

* Edit the containerd config

```shell
vi /etc/containerd/config.toml
```

* Add the following lines under the main plugins section

```toml
[plugins]
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runsc]
runtime_type = "io.containerd.runsc.v1"
...
```

* Restart containerd service (your node-shell will be closed and you have to reconnect)

```shell
systemctl restart containerd
```

## Create a runtimeClass

```shell
kubectl apply -f runtimeclass.yaml
```

## Run a Pod that makes use of the RuntimeClass

```shell
kubectl -n ${YOURNAME} apply -f gvisor-pod.yaml
```

---

## Verify

* Get the kernel version of the gvisor Pod

```shell
kubectl -n ${YOURNAME} exec -it gvisorpod -- uname -a
Linux gvisorpod 4.4.0 #1 SMP Sun Jan 10 15:06:54 PST 2016 x86_64 GNU/Linux
```

* Compare the kernel version with a standard Pod

```shell
kubectl -n ${YOURNAME} exec -it <PODNAME> -- uname -a 
5.15.0-87-generic #97-Ubuntu SMP Mon Oct 2 21:09:21 UTC 2023 x86_64 Linux
```

---

## Conclusion

* Using a runtimeClass with a runtime handler can increase security of your cluster.
* Depending on the runtime used it might reduce performance.

---

## Reference

* gvisor documentation

https://gvisor.dev/docs/user_guide/install/#install-latest

https://gvisor.dev/docs/user_guide/containerd/quick_start/

* RuntimeClass [documentation](https://kubernetes.io/docs/concepts/containers/runtime-class/)
