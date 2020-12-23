package main

import (
	"sigs.k8s.io/kustomize/api/k8sdeps/kunstruct"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// constants
const defaultConfigString = `fieldprefix: rnd
fieldSpecs:
  - path: metadata/namespace
    create: true`

func newPluginHelpers(resmapFactory *resmap.Factory) *resmap.PluginHelpers {
	return resmap.NewPluginHelpers(nil, nil, resmapFactory)
}

func newResMapFactory() *resmap.Factory {
	resourceFactory := resource.NewFactory(kunstruct.NewKunstructuredFactoryImpl())
	return resmap.NewFactory(resourceFactory, nil)
}

func getDefaultConfig() (transformer, error) {
	var defaultConfig transformer
	err := yaml.Unmarshal([]byte(defaultConfigString), &defaultConfig)
	return defaultConfig, err
}
