# PodDisruptionBudget (pdb)

## Preparation

This hands-on exercise continues the previous lecture [120a-hpa](../120a-hpa/README.md)
Only proceed if you have accomplished the previous section!

---

## Exercise

* Apply PodDisruptionBudget resource

  ```shell
  kubectl apply -f php-apache-pdb.yaml
  ```

* Delete HPA, scale down to one replica and find out what node the last pod is runnning on
  
  ```shell
  kubectl delete -f ../120a-hpa/php-apache-hpa.yaml
  kubectl scale -f ../120a-hpa/php-apache-deployment.yaml --replicas=1
  kubectl delete pod load-generator
  kubectl get pod -o wide
  kubectl get node
  ```

* Try to drain node (this should only be done by **one participant**)
  
  ```shell
  kubectl drain "NODE-NAME" --ignore-daemonsets --delete-emptydir-data # extra options because of metakube specific pods running
  ```

* you should see a message like this:

  ```shell
  error when evicting pods/"php-apache-f966667ff-wshts" -n "YOURNAME" (will retry after 5s): 
  Cannot evict pod as it would violate the pod's disruption budget.
  ```

* Stop our experiment and let the node join again

  ```shell
  kubectl uncordon "NODE-NAME"
  ```

## Conclusion

* The PodDisruptionBudget prevents the pod from being evicted.

---

## Clean up

* Tear down everything

  ```shell
  kubectl delete -f ../120a-hpa/php-apache-hpa.yaml \
    -f ../120a-hpa/php-apache-service.yaml \
    -f ../120a-hpa/php-apache-deployment.yaml \
    -f php-apache-pdb.yaml
  ```
