package main

import (
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/filtersutil"
	"sigs.k8s.io/yaml"
)

// configuration of the random field generator
// the random prefix is something that will guide the name generation
type transformer struct {
	FieldPrefix string            `json:"fieldprefix,omitempty" yaml:"fieldprefix,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

// GlobalPlugin used in other parts of the module
var GlobalPlugin transformer

func (tr *transformer) Config(_ *resmap.PluginHelpers, c []byte) error {
	tr.FieldPrefix = ""
	tr.FieldSpecs = nil
	return yaml.Unmarshal(c, tr)
}

func (tr *transformer) Transform(m resmap.ResMap) error {
	suffix := rgInstance.suffix()
	for _, r := range m.Resources() {
		if r.IsEmpty() {
			// Don't mutate empty objects?
			continue
		}

		filter := Filter{
			FieldPrefix: tr.FieldPrefix,
			Suffix:      suffix,
			FieldSpecs:  tr.FieldSpecs,
		}

		err := filtersutil.ApplyToJSON(filter, r)
		if err != nil {
			return err
		}
	}
	return nil
}
