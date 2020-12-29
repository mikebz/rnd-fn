This is an example of setting a random namespace.  In this example
the instrucitons and the cluster configuration are all in one file.

To make the execution easier we are using --enable-exec flag that allows
us to use the rnd-fn binary post built.

Example usage in the shell:
```sh
kpt fn run . --enable-exec
```
This will add the suffix to the namespaces in all objects.