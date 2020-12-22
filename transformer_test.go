package main

import (
	"fmt"
	"testing"
)

// test setup function that creates the transformer and resmap
func runNamespaceTransformerE(config, input string) (string, error) {
	resmapFactory := newResMapFactory()
	resMap, err := resmapFactory.NewResMapFromBytes([]byte(input))
	if err != nil {
		return "", err
	}

	var plugin *plugin = &GlobalPlugin
	err = plugin.Config(nil, []byte(config))
	if err != nil {
		return "", err
	}
	defaultConfig, err := getDefaultConfig()
	if err != nil {
		return "", err
	}
	if len(plugin.FieldSpecs) == 0 {
		plugin.FieldSpecs = defaultConfig
	}
	err = plugin.Transform(resMap)
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

func TestSimpleNamespace(t *testing.T) {
	config := `
fieldprefix: test
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
  publishNotReadyAddresses: true
  ports:
  - port: 2380
    name: etcd-server-ssl
`

	expected := `apiVersion: v1
kind: Service
metadata:
  name: the-service
  namespace: unique-ns-1231231
spec:
  clusterIP: None
  publishNotReadyAddresses: true
  ports:
  - port: 2380
    name: etcd-server-ssl
`

	output := runNamespaceTransformer(t, config, input)
	if output != expected {
		fmt.Println("Actual:")
		fmt.Println(output)
		fmt.Println("===")
		fmt.Println("Expected:")
		fmt.Println(expected)
		t.Fatalf("Actual doesn't equal to expected")
	}
}
