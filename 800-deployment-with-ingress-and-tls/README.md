# Deployment with Ingress and TLS

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

Deploy a full web application to your namespace which uses external-dns, ingress-nginx and cert-manager.

* First, adjust your individual hostname in `web-application/deployment/ingress.yaml`

* Deploy application

  ```shell
  kubectl apply -f web-application/deployment/web-application.yaml
  ```

* Deploy ingress resource

  ```shell
  kubectl apply -f web-application/deployment/ingress.yaml
  ```

## Conclusion

* we learned how powerful the combined tools external-dns, ingress-nginx and cert-manager are
* it is easy to publish applications and get certificates and a DNS name directly from a k8s resource
