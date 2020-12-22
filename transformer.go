package main

import (
	"sigs.k8s.io/kustomize/api/filters/prefixsuffix"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/filtersutil"
	"sigs.k8s.io/yaml"
)

// configuration of the random field generator
// the random prefix is something that will guide the name generation
type plugin struct {
	FieldPrefix string            `json:"fieldprefix,omitempty" yaml:"fieldprefix,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

// GlobalPlugin used in other parts of the module
var GlobalPlugin plugin

func (p *plugin) Config(_ *resmap.PluginHelpers, c []byte) (err error) {
	p.FieldPrefix = ""
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *plugin) Transform(m resmap.ResMap) error {
	for _, r := range m.Resources() {
		if r.IsEmpty() {
			// Don't mutate empty objects?
			continue
		}

		filter := prefixsuffix.Filter{
			Suffix:    "-" + rgInstance.suffix(),
			FieldSpec: p.FieldSpecs[0], // TODO: create a filter with multiple fieldspecs
		}

		err := filtersutil.ApplyToJSON(filter, r)
		if err != nil {
			return err
		}
	}
	return nil
}
