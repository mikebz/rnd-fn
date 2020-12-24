package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// basic test that makes sure that our default
// config indeed parses into somthing that we expect
// that is a prefix and field specs
func TestSimpleConfig(t *testing.T) {
	resmapFactory := newResMapFactory()
	pluginHelpers := newPluginHelpers(resmapFactory)

	tr := transformer{}
	err := tr.Config(pluginHelpers, []byte(defaultConfigString))
	assert.NoError(t, err)
	assert.NotNil(t, tr.FieldPrefix)
	assert.NotNil(t, tr.FieldSpecs)
	assert.Len(t, tr.FieldSpecs, 1)
}

// test that verifies that multipel field specs
// are parsed in to multiple FieldSpec objects
func TestMultipleSpecValues(t *testing.T) {

	config := `---
fieldprefix: rnd
fieldSpecs:
  - path: metadata/namespace
    group: foo
    kind: Bar
    apiVersion: foo/v1beta1
    create: true
`

	resmapFactory := newResMapFactory()
	pluginHelpers := newPluginHelpers(resmapFactory)

	tr := transformer{}
	err := tr.Config(pluginHelpers, []byte(config))
	assert.NoError(t, err)
	assert.NotNil(t, tr.FieldPrefix)
	assert.NotNil(t, tr.FieldSpecs)
	assert.Len(t, tr.FieldSpecs, 1)
}

func TestMultipleSpecs(t *testing.T) {

	config := `---
fieldprefix: rnd
fieldSpecs:
  - path: metadata/namespace
    create: true
  - path: metadata/label
    create: true
`

	resmapFactory := newResMapFactory()
	pluginHelpers := newPluginHelpers(resmapFactory)

	tr := transformer{}
	err := tr.Config(pluginHelpers, []byte(config))
	assert.NoError(t, err)
	assert.NotNil(t, tr.FieldPrefix)
	assert.NotNil(t, tr.FieldSpecs)
	assert.Len(t, tr.FieldSpecs, 2)
}
