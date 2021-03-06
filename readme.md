# rnd-fn kpt function

The goal of this function is to enable the use case where you'd like to 
randomize a particular element - maybe a label or a namespace for a kpt package.

An real world example of where this can be handy is creating a dev/test
environment for a branch or feature and then tearing it down later.  Multiple
teams in your organization can be creating these environments and the 
blueprint for the environment can be the same application, but the unique
namespace allows you to identify which is which and possibly clean them up.

## Usage documentation

Configured using a ConfigMap with the following keys:

fieldprefix: the prefix to look for..
fieldSpecs: A list of specification to select the resources and fields that 
the randomly generated value will be applied to.

Example:

To add a random number suffix to the namespace 'foobar' so it becomes 'foobar-3143153' 
 to all resources, use this example:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
  namespace: foobar
data:
  fieldprefix: foobar
```

You can use key 'fieldSpecs' to specify the resource selector you
want to use. By default, the function will use this field spec:

```
- path: metadata/namespace
  create: true
```

This means a 'metadata/namespace' field will be added to all resources
with namespaceable kinds. Whether a resource is namespaceable is determined
by the Kubernetes API schema. If the API path for that kind contains
'namespaces/{namespace}' then the resource is considered namespaceable. Otherwise
it's not.

For more information about API schema used in this function, please take a look at
https://github.com/kubernetes-sigs/kustomize/tree/master/kyaml/openapi

Field spec has following fields:

- group: Select the resources by API version group. Will select all groups
	if omitted.
- version: Select the resources by API version. Will select all versions
	if omitted.
- kind: Select the resources by resource kind. Will select all kinds
	if omitted.
- path: Specify the path to the field that the value will be updated. This field
	is required.
- create: If it's set to true, the field specified will be created if it doesn't
	exist. Otherwise the function will only update the existing field.

Example:

To add a namespace 'foobar-23142425' to Deployment resource only:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
data:
  fieldprefix: foobar
  fieldSpecs:
    - path: metadata/namespace
      kind: Deployment
      create: true
```

## End to end user journey.

Before you start this assumes that you have built and installed the rnd-fn binary using `go install .`

Create a directory in which you will host your config, but you can also use an existing repo
or existing directory.
```
> mkdir rnd-usecase
> cd rnd-usecase
```

fetch the sample package into the directory
```
> kpt pkg get https://github.com/mikebz/rnd-example .
fetching package / from https://github.com/mikebz/rnd-example to rnd-example
```

Now we can run the function in the current folder and see the function randomize the names
```
> kpt fn run --enable-exec rnd-example/selenium --fn-path rnd-example/fn-config 
```

You can now deploy it with `kpt live apply`, ConfigSync or directly with `kubectl apply`

The effect on one of the files is as follows.  `selenium-hub-svc.yml` before: 

```
apiVersion: v1
kind: Service
metadata:
  name: selenium-hub
  labels:
    app: selenium-hub
spec:
  ports:
  - port: 4444
    targetPort: 4444
    name: port0
  selector:
    app: selenium-hub
  type: NodePort
  sessionAffinity: None
```

after:
```
apiVersion: v1
kind: Service
metadata:
  labels:
    app: selenium-hub-8826718
  name: selenium-hub-8826718
spec:
  ports:
  - name: port0
    port: 4444
    targetPort: 4444
  selector:
    app: selenium-hub-8826718
  sessionAffinity: None
  type: NodePort
  ```

  The function sees a field named with a particular patterns like `fieldname-#######` and realizes that the field has been randomized already and doesn't append any more suffixes to it.