# Resources

## Requests and Limits

## Preparation

* Before you begin with the actual exercise please make sure to follow these steps to work in your own environment:

  ```shell
  read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
  export YOURNAME
  kubectl create ns ${YOURNAME}
  kubectl label namespace ${YOURNAME} deepdive-pgd=true
  kubectl config set-context --current --namespace=${YOURNAME}
  ```

* Clone this repository to your working station and change into the directory for the following exercises

---

## Exercise

### Part 1 - Resource requests and limits

* Create an nginx deployment with requests and limits set.
* Verify if the pod is running.

  ```shell
  kubectl apply -f deploy-requests-limits.yaml
  kubectl get po -l app=nginx
  ```

* Result: the pod is in `CrashLoopBackoff` state because memory limits are to low.
* Increase the memory limit to `50Mi` in the file `deploy-requests-limits.yaml` and recreate the deployment.

```shell
  kubectl apply -f deploy-requests-limits.yaml
  kubectl get po -l app=nginx
  ```

* Result: Is the pod starting up now?

### Clean up

* Delete the deployment

  ```shell
  kubectl delete -f deploy-requests-limits.yaml
  ```

---

### Part 2 - Namespace limit range

* Create a `LimitRange` in your namespace

  ```shell
  kubectl -n ${YOURNAME} apply -f limit-range.yaml
  ```

* Show your settings

  ```shell
  kubectl describe limitrange mem-limit-range
  ```

* Verify the LimitRange is enforced

* Create a pod without resource settings in your namespace

  ```shell
  kubectl run nginx-no-limits --image=nginx:1.23.0
  ```

* Verify it gets the settings from the LimitRange assigned automatically

  ```shell
  kubectl get po nginx-no-limits -o yaml
  ```

* Look for the `LimitRanger` annotation
* Look for the `Resources` section

---

### Additional commands for a better resource overview

* show all resources of a node available for workloads

  ```shell
  kubectl get nodes -o yaml | grep allocatable -A5
  kubectl describe nodes | grep <POD-NAME>
  ```

---

### Clean up

* Delete the pod again

  ```shell
  kubectl delete po nginx-no-limits
  ```

---

### Conclusion

- It is highly recommended to set resource requests and limits to a container
- These settings are also required by a HorizontalPodAutoscaler we will deploy in another hands-on session
- You can avoid Pods lacking requests and limit completely by applying a LimitRanger in the given namespace
