This is an example of setting a random namespace.  The function config
is in a separate file `fn-config.yml`.

To make the execution easier we are using --enable-exec flag that allows
us to use the rnd-fn binary post built.

Example usage in the shell:
```sh
kpt fn run . --enable-exec
```

This will add the suffix to the namespaces in all objects.