package main

import (
	"github.com/stretchr/testify/assert"
	filtertest_test "sigs.k8s.io/kustomize/api/testutils/filtertest"
	"sigs.k8s.io/kustomize/api/types"
	"testing"
)

func TestSimpleSet(t *testing.T) {
	input := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: instance
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: instance
`
	expected := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: instance-123
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: instance-123
`
	filter := Filter{
		FieldPrefix: "instance",
		Suffix:      "123",
		FieldSpecs:  []types.FieldSpec{{Path: "metadata/name"}},
	}

	actual := filtertest_test.RunFilter(t, input, filter)

	assert.Equal(t, expected, actual)
}

func TestRepeatedSet(t *testing.T) {
	input := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: instance-232
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: instance-232
`
	expected := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: instance-232
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: instance-232
`
	filter := Filter{
		FieldPrefix: "instance",
		Suffix:      "123",
		FieldSpecs:  []types.FieldSpec{{Path: "metadata/name"}},
	}

	actual := filtertest_test.RunFilter(t, input, filter)

	assert.Equal(t, expected, actual)
}

func TestRegex(t *testing.T) {
	input := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: foo-area-51
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: bar-area-51
`
	expected := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: foo-area-51-232
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: bar-area-51-232
`
	filter := Filter{
		FieldPrefix: "\\w*-area-\\d*",
		Suffix:      "232",
		FieldSpecs:  []types.FieldSpec{{Path: "metadata/name"}},
	}

	actual := filtertest_test.RunFilter(t, input, filter)

	assert.Equal(t, expected, actual)
}

func TestRegexRepeatedSet(t *testing.T) {
	input := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: foo-area-51-112
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: bar-area-51-112
`
	expected := `apiVersion: example.com/v1
kind: Foo
metadata:
  name: foo-area-51-112
---
apiVersion: example.com/v1
kind: Bar
metadata:
  name: bar-area-51-112
`
	filter := Filter{
		FieldPrefix: "\\w*-area-\\d*",
		Suffix:      "232",
		FieldSpecs:  []types.FieldSpec{{Path: "metadata/name"}},
	}

	actual := filtertest_test.RunFilter(t, input, filter)

	assert.Equal(t, expected, actual)
}
