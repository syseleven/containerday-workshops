# Probes

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

## Exercise

## Readiness and Liveness Probes

* Create an nginx deployment with `readinessProbe` and `livenessProbe` set.
* Verify if the pod is running.

  ```shell
  kubectl apply -f deploy-probes.yaml
  kubectl get po -l app=nginx-probes
  ```

### Clean up

* Delete the deployment

  ```shell
  kubectl delete -f deploy-probes.yaml
  ```
