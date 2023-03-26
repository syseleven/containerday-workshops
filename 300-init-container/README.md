# Init Container

---

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

Deploy an application with init containers ensuring the application finds a service it depends on.

### Deploy application

* Lets create a Pod with three containers. The first container is our main container and the second and third are init containers that wait until some service will show up

```shell
kubectl apply -f myapp.yaml

kubectl get -f myapp.yaml
```

* Lets see what our Pod is doing

```shell
kubectl describe -f myapp.yaml

kubectl logs myapp-pod -c init-myservice # Inspect the first init container

kubectl logs myapp-pod -c init-mydb      # Inspect the second init containers
```

### Deploy a service

* Now we will create the missing services so that our init containers can finally start

```shell
kubectl apply -f services.yaml
```

### Verify

* Check that the pod is now running

```shell
kubectl get -f myapp.yaml
```

### Clean up

* Lets tear down everything

```shell
kubectl delete -f myapp.yaml -f services.yaml
```
