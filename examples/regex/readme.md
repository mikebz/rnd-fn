This is an example of setting names and namespaces
with a regular expression.  The fieldspecs and regular expression are in `fn-config.yml`.

To make the execution easier we are using --enable-exec flag that allows
us to use the rnd-fn binary post built.

Example usage in the shell:
```sh
kpt fn run . --enable-exec
```

This will add the suffix to the namespaces in all objects.