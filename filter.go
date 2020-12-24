package main

import (
	"fmt"
	"regexp"
	"sigs.k8s.io/kustomize/api/filters/filtersutil"
	"sigs.k8s.io/kustomize/api/filters/fsslice"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// Filter applies the Suffix to the resources that match the prefix
type Filter struct {
	FieldPrefix string            `json:"fieldprefix,omitempty" yaml:"fieldprefix,omitempty"`
	Suffix      string            `json:"suffix,omitempty" yaml:"suffix,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpec,omitempty" yaml:"fieldSpec,omitempty"`
}

// Filter goes through the list of RNodes and applies the filter to them.
func (f Filter) Filter(nodes []*yaml.RNode) ([]*yaml.RNode, error) {
	return kio.FilterAll(yaml.FilterFunc(f.run)).Filter(nodes)
}

func (f Filter) run(node *yaml.RNode) (*yaml.RNode, error) {
	err := node.PipeE(fsslice.Filter{
		FsSlice:    f.FieldSpecs,
		SetValue:   f.evaluateField,
		CreateKind: yaml.ScalarNode,
		CreateTag:  yaml.NodeTagString,
	})
	return node, err
}

// this is a function that gets called after
// all of the field specs are matched from the fieldspec slices
func (f Filter) evaluateField(node *yaml.RNode) error {
	currentValue := node.YNode().Value

	// we need to make sure that the prefix is matched
	// and that the value has not been set yet.
	matchPrefix, err := regexp.MatchString(f.FieldPrefix, currentValue)
	if err != nil {
		return err
	}

	matchSet, err := regexp.MatchString(f.FieldPrefix+"-\\d*", currentValue)
	if err != nil {
		return err
	}

	if matchPrefix && !matchSet {
		newValue := fmt.Sprintf("%s-%s", currentValue, f.Suffix)
		return filtersutil.SetScalar(newValue)(node)
	}

	return nil
}
