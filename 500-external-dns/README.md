# External DNS

## Task

Deploy external-dns in the cluster as a central service.

**This installation is required only once per cluster. (Only 1 participant!)**

[ArtifactHub - external-dns](https://artifacthub.io/packages/helm/bitnami/external-dns)

  ```shell
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update
  ```

* Replace all occurrences of `CHANGEME` in `values.yaml` with the secrets
provided by the trainer.

* Install external-dns with Helm:

  ```shell
  kubectl create namespace external-dns
  ```

  ```shell
  helm upgrade --install external-dns --namespace external-dns --values values.yaml bitnami/external-dns --version=6.19.1
  ```

* Check if ExternalDNS Pod is up and running

  ```shell
  kubectl --namespace=external-dns get pods -l "app.kubernetes.io/name=external-dns,app.kubernetes.io/instance=external-dns"
  ```

* Check ExternalDNS Pod Logs

  ```shell
  kubectl --namespace=external-dns logs -l "app.kubernetes.io/name=external-dns,app.kubernetes.io/instance=external-dns"
  ```

* Deploy nginx test with dns entry:

  ```shell
  kubectl apply -f externaldns-test.yaml --namespace external-dns
  ```

* Check Deployment, Services and Pods from previous Deployment and watch out for External IP for Service `service/nginx`

  ```shell
  kubectl get deployment,pods,services --namespace external-dns
  ```

* Should look like this (LoadBalancer service might take a few minutes to create)

  ```shell
  NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP      PORT(S)        AGE
  service/nginx          LoadBalancer   10.240.26.62    185.56.130.x   80:32428/TCP   2m19s
  ```

* View the logs of external-dns

  ```shell
  kubectl -n external-dns logs -f -l app.kubernetes.io/name=external-dns
  ```
  
  Expected result:

  ```shell
  time="2023-11-13T08:48:56Z" level=info msg="All records are already up to date"
  time="2023-11-13T08:49:56Z" level=info msg="Applying provider record filter for domains: [metakube.org. .metakube.org.]"
  ```

### Verify

* Visit URL https://nginx.workshop.metakube.org

---

### Clean up

* Delete the externaldns-test pod

  ```shell
  kubectl delete -f externaldns-test.yaml --namespace external-dns
  ```
