# Horizontal Pod Autoscaler (HPA)

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

## Part 1 - Create deployment

Task: Create a deployment which contains elements from the previous lectures.
Additionally we deploy a `horizontalPodAutoscaler` to scale pods up and down dynamically.

* Apply deployment and check if everything works as expected

  ```shell
  kubectl apply -f php-apache-deployment.yaml
  
  kubectl get -f php-apache-deployment.yaml
  kubectl get pod,deploy
  kubectl describe deployment php-apache
  ```

* We need a service, with a clusterIP, so that the pods respond to our requests 

  ```shell
  kubectl apply -f php-apache-service.yaml
  ```

* Check if service is up as expected

  ```shell
  kubectl get -f php-apache-service.yaml
  kubectl get svc php-apache -o wide
  ```

* Look out for Endpoints. These are your Pods!
  
  ```shell
  kubectl describe svc php-apache
  ```

  ```shell
  kubectl get pods -o wide
  ```

* Scale down the deployment

  ```shell
  kubectl scale -f php-apache-deployment.yaml --replicas=1
  kubectl get -f php-apache-deployment.yaml
  kubectl get pods
  ```

## Part 2 - Deploy the HorizontalPodAutoscaler

* Let's check the content of the HorizontalPodAutoscaler in `php-apache-hpa.yaml`

  ```yaml
  apiVersion: autoscaling/v1
  kind: HorizontalPodAutoscaler
  metadata:
    name: php-apache
  spec:
    maxReplicas: 10
    minReplicas: 1
    scaleTargetRef:
      apiVersion: apps/v1
      kind: Deployment
      name: php-apache
    # das Target bezieht sich auf den CPU Request des gesamten Pods
    # daher ist es erforderlich resource requests zu setzen
    targetCPUUtilizationPercentage: 25
  ```

* Apply the `HorizontalPodAutoscaler`

  ```shell
  kubectl apply -f php-apache-hpa.yaml
  kubectl get -f php-apache-hpa.yaml
  ```

## Part 3 - Create load on the pods

* Start a new pod with the busybox image which we will use to stress our service.

* We have to be a little patient until the results shows up

  ```shell
  kubectl run -i --tty load-generator --image=busybox -- /bin/sh
  
  # reconnect to the pod if you want to use it again after you exit it
  kubectl attach load-generator -c load-generator -i -t
  
  # create some load on the pods
  # replace <YOURNAME> with you data
  while true; do wget -q -O- http://php-apache.<YOURNAME>.svc.cluster.local; done
  ```

---

Optional:

* Alternative to busybox image via port-forward
  * Lets the local machine forward to the cluster.
  * Either to a pod directly or to a service.
  * To test pods, Bret Fisher's shpod.yml is also useful.
  * You need to run the scenario with two terminal windows.

  ```shell
  kubectl port-forward svc/php-apache 8080:80
  kubectl logs -f -l app=php-apache --all-containers
  ```

---

* Lets inspect what's going on

  ```shell
  kubectl get hpa
  kubectl top pod
  kubectl get deployment php-apache
  ```

---

## Conclusion

* So far we deployed a HPA and it performed a pod scale-up under increasing load.
* Scaling criteria is a percentage of the pod's CPU utilization.

## Proceed with the next lecture

Do not delete any of the obove resources and continue with the next exercise
about [PodDisruptionBudgets](../120b-pdb/README.md) (PDB)
