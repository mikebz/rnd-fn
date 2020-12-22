package main

const usage = `random kpt function

Configured using a ConfigMap with the following keys:

fieldprefix: the prefix to look for..
fieldSpecs: A list of specification to select the resources and fields that 
the randomly generated value will be applied to.

Example:

To add a suffix to the namespace 'foobar' so it becomes 'foobar-3143153' 
(or any other number) to all resources:

apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
  namespace: foobar
data:
  fieldprefix: foobar

You can use key 'fieldSpecs' to specify the resource selector you
want to use. By default, the function will use this field spec:

- path: metadata/namespace
  create: true

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

To add a namespace 'foobar' to Deployment resource only:

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

`
