# nginx Ingress Controller

**This installation is required only once per cluster.**

[ArtifactHub - ingress-nginx](https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx)

* Install ingress controller "ingress-nginx"

* Create a namespace

  ```shell
  kubectl create namespace ingress-nginx
  ```

* Add ingress-nginx Helm Repo and update Cache

  ```shell
  helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
  helm repo update
  ```

* Install ingress-nginx with provided values.yaml

  ```shell
  helm upgrade --install -f values.yaml --namespace ingress-nginx ingress-nginx ingress-nginx/ingress-nginx --version=4.6.0
  ```

---

* **Note!** Please be a little patient since we deploy a DaemonSet on all Nodes!
* Also creating the LoadBalancer service might take a few minutes!

---

### Verify
* Check if ingress-nginx is up and running

  ```shell
  kubectl --namespace ingress-nginx get services -o wide -w ingress-nginx-controller
  ```
