# Helmfile

**This is just a dry-run example - not intended to install actually.**

Helmfile is a tool to combine multiple Helm Charts and their respective installation commands.

## Preparation

* Helmfile project [on GitHub](https://github.com/helmfile/helmfile)

* Install the tool `helmfile` following its [documentation](https://github.com/helmfile/helmfile/tree/main#installation)

* Verify installation

  ```shell
  $ helmfile --version
  helmfile version 0.158.1
  ```

---

## Task

Combine all required resources of ingress-nginx, external-dns and cert-manager for a unified installation file.

* Inspect the `helmfile.yaml`
  * There are only 2 values files, but 3 releases. Discuss.
  * Note the `needs` section.
* Use helmfile to build its manifests locally for closer inspection.

  ```shell
  helmfile template --output-dir .
  ```

* Use helmfile to show diffs regarding the workshop cluster.
  * What are the differences?

  ```shell
  helmfile diff
  ```

---

## Conclusion

It stores all installation instructions in one single file which can be used for documentation and IaaC.

It creates another layer of abstraction.

Key features:
- combine Helm Charts from different places (local, registry)
- inject multiple values files
- create dependencies between the releases