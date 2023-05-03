# cert-manager

**This installation is required only once per cluster.**

[ArtifactHub - cert-manager](https://artifacthub.io/packages/helm/cert-manager/cert-manager)

## Install and configure cert-manager

  ```shell
  helm repo add jetstack https://charts.jetstack.io
  ```

  ```shell
  helm repo update
  ```

  ```shell
  kubectl create namespace cert-manager
  ```

  ```shell
  helm upgrade --install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.11.1 --set installCRDs=true
  ```

## Configuration for route53

* Add AWS cert-manager Secret Key to `credentials-secret.yaml`
* Add AWS cert-manager Access Key to `clusterissuer.yaml`

* Create ClusterIssuer

  ```shell
  kubectl apply -f credentials-secret.yaml
  kubectl apply -f clusterissuer.yaml
  ```

* Create wildcard certificate

  ```shell
  kubectl apply -f certificate.yaml
  ```

* Last step: activate in NGINX [values.yaml](../600-ingress-nginx/values.yaml)

  ```yaml
  controller:
    # Activate after successful cert-manager installation
    extraArgs:
      default-ssl-certificate: ingress-nginx/workshop-tls-cert 
  ```

* Upgrade ingress-nginx helm release

  ```shell
  helm upgrade --install -f ../600-ingress-nginx/values.yaml --namespace ingress-nginx ingress-nginx ingress-nginx/ingress-nginx --version=4.3.0
  ```
