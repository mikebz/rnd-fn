package main

import (
	"fmt"
	"testing"
)

type randomMock struct{}

func (rg randomMock) suffix() string {
	return "1231231"
}

func (rg randomMock) value(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, rg.suffix())
}

// test setup function that creates the transformer and resmap
func runNamespaceTransformerE(config, input string) (string, error) {

	rgInstance = randomMock{}

	resmapFactory := newResMapFactory()
	resMap, err := resmapFactory.NewResMapFromBytes([]byte(input))
	if err != nil {
		return "", err
	}

	var testTr *transformer = &GlobalPlugin
	err = testTr.Config(nil, []byte(config))
	if err != nil {
		return "", err
	}
	err = testTr.Transform(resMap)
	if err != nil {
		return "", err
	}
	y, err := resMap.AsYaml()
	if err != nil {
		return "", err
	}
	return string(y), nil
}

// test helper that returns the result of the transformation
func runNamespaceTransformer(t *testing.T, config, input string) string {
	s, err := runNamespaceTransformerE(config, input)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

// testing the simple namespace transformation
func TestSimpleNamespace(t *testing.T) {
	config := `
fieldprefix: unique-ns
fieldSpecs:
- path: metadata/namespace
`

	input := `apiVersion: v1
kind: Service
metadata:
  name: the-service
  namespace: unique-ns
spec:
  clusterIP: None
  ports:
  - port: 2380
  publishNotReadyAddresses: true
`

	expected := `apiVersion: v1
kind: Service
metadata:
  name: the-service
  namespace: unique-ns-1231231
spec:
  clusterIP: None
  ports:
  - port: 2380
  publishNotReadyAddresses: true
`

	actual := runNamespaceTransformer(t, config, input)
	if actual != expected {
		fmt.Println("Actual:")
		fmt.Println(actual)
		fmt.Println("===")
		fmt.Println("Expected:")
		fmt.Println(expected)
		t.Fatalf("Actual doesn't equal to expected")
	}
}

// testing the simple lable transformation.
func TestSimpleLabel(t *testing.T) {
	config := `
fieldprefix: test
fieldSpecs:
- path: metadata/label
`

	input := `apiVersion: v1
kind: Service
metadata:
  label: test
  name: the-service
  namespace: my-ns
spec:
  clusterIP: None
  ports:
  - port: 2380
  publishNotReadyAddresses: true
`

	expected := `apiVersion: v1
kind: Service
metadata:
  label: test-1231231
  name: the-service
  namespace: my-ns
spec:
  clusterIP: None
  ports:
  - port: 2380
  publishNotReadyAddresses: true
`

	actual := runNamespaceTransformer(t, config, input)
	if actual != expected {
		fmt.Println("Actual:")
		fmt.Println(actual)
		fmt.Println("===")
		fmt.Println("Expected:")
		fmt.Println(expected)
		t.Fatalf("Actual doesn't equal to expected")
	}
}

// testing the transform of a regular expression in a fieldPrefix.
func TestTransformRegex(t *testing.T) {
	config := `
fieldprefix: "unique-\\w+"
fieldSpecs:
- path: metadata/name
- path: metadata/namespace
`

	input := `apiVersion: v1
kind: ConfigMap
metadata:
  name: unique-cm
  namespace: unique-ns
`

	expected := `apiVersion: v1
kind: ConfigMap
metadata:
  name: unique-cm-1231231
  namespace: unique-ns-1231231
`

	actual := runNamespaceTransformer(t, config, input)
	if actual != expected {
		fmt.Println("Actual:")
		fmt.Println(actual)
		fmt.Println("===")
		fmt.Println("Expected:")
		fmt.Println(expected)
		t.Fatalf("Actual doesn't equal to expected")
	}

}
