# Canary Deployments

**This installation is required by every participant.**

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

## Task

Create a "blue" and "green" deployment. Shift the traffic over to "green" continuously.  

## Adjust

* Adjust your individual hostname in `web-application-blue/ingress-blue.yaml`

* Adjust your individual hostname in `web-application-green/ingress-green-canary.yaml`

## Blue

* Deploy blue application, service and ingress resource

  ```shell
  kubectl apply -f web-application-blue/
  ```

### Verify

* Check if the blue application is available

  ```shell
  # check k8s resources
  kubectl get po,svc,ing
  
  # send requests to our blue application
  for i in {1..10} ; do curl https://web-application-canary-${YOURNAME}.workshop.metakube.org/echo ; done
  ```

### Result

```shell
"echo-blue"
"echo-blue"
"echo-blue"
"echo-blue"
[...]
"echo-blue"
```

---

## Green (Canary)

* Deploy green application, service and ingress resource

  ```shell
  kubectl apply -f web-application-green/
  ```

### Verify

* Check if the blue application is available

  ```shell
  # check k8s resources
  kubectl get po,svc,ing
  
  # send requests to our blue application
  for i in {1..10} ; do curl https://web-application-canary-${YOURNAME}.workshop.metakube.org/echo ; done
  ```

### Result

```shell
"echo-blue"
"echo-green"
"echo-blue"
"echo-green"
[...]
"echo-blue"
```

---

## Increase traffic routed to green

* Increase the weight value in the file `web-application-green/ingress-green-canary.yaml` in steps like 50, 75, 100

  ```yaml
  [...]
  annotations:
    nginx.ingress.kubernetes.io/canary-weight: "75" # <-- please adjust later up to 100
  [...]
  ```

* After each increase update the ingress resource in the cluster:

  ```shell
  kubectl apply -f web-application-green/ingress-green-canary.yaml
  ```

* Check if the traffic shifts according to your weight setting:

  ```shell
  # send requests to our blue application
  for i in {1..10} ; do curl https://web-application-canary-${YOURNAME}.workshop.metakube.org/echo ; done
  ```

### Result

```shell
"echo-green"
"echo-green"
"echo-green"
"echo-green"
[...]
"echo-blue"
```

---

## Conclusion

With blue/green deployments (Canary deployments) it is possible to introduce a new application version
and use the canary mechanism of ingress-nginx to continuously shift traffic over to the new target.

After a successful migration 100% of the traffic can be routed to the green deployment and the blue deployment
may be removed seamlessly.
