# Security

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

## Part 1 - Pod and Container Security

* Deploy a secure pod

  ```shell
  kubectl -n ${YOURNAME} apply -f XXXXXXXXXXXXXXX
  ```

---

## Part 2 - Pod Security Standards

* Configure a namespace to use "PSS"

  ```shell
  kubectl label --overwrite ns ${YOURNAME} \
  pod-security.kubernetes.io/warn=baseline \
  pod-security.kubernetes.io/warn-version=latest
  ```

* Now deploy a pod violating "PSS"

  ```shell
  kubectl -n ${YOURNAME} apply -f XXXXXXXXXXXXXXX
  ```

---

## Cleanup

* remove PSS from you namespace again

  ```shell
  kubectl label --overwrite ns ${YOURNAME} \
  pod-security.kubernetes.io/warn- \
  pod-security.kubernetes.io/warn-version-
  ```

---

### Conclusion

A deployment strategy of `Recreate` causes a downtime for the application.
