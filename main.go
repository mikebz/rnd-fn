// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	var tr *transformer = &GlobalPlugin
	defaultConfig, err := getDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	resourceList := &framework.ResourceList{}
	resourceList.FunctionConfig = map[string]interface{}{}
	resmapFactory := newResMapFactory()
	pluginHelpers := newPluginHelpers(resmapFactory)

	cmd := framework.Command(resourceList, func() error {
		resMap, err := resmapFactory.NewResMapFromRNodeSlice(resourceList.Items)
		if err != nil {
			return errors.Wrap(err, "failed to convert items to resource map")
		}
		dataField, err := getDataFromFunctionConfig(resourceList.FunctionConfig)
		if err != nil {
			return errors.Wrap(err, "failed to get data field from function config")
		}
		dataValue, err := yaml.Marshal(dataField)
		if err != nil {
			return errors.Wrap(err, "error when marshal data values")
		}

		err = tr.Config(pluginHelpers, dataValue)
		if err != nil {
			return errors.Wrap(err, "failed to config plugin")
		}

		// if we aer missing some elements of config let's
		// get the defaults
		if len(tr.FieldPrefix) == 0 {
			tr.FieldPrefix = defaultConfig.FieldPrefix
		}
		if len(tr.FieldSpecs) == 0 {
			tr.FieldSpecs = defaultConfig.FieldSpecs
		}

		err = tr.Transform(resMap)
		if err != nil {
			return errors.Wrap(err, "failed to run transformer")
		}

		resourceList.Items, err = resMap.ToRNodeSlice()
		if err != nil {
			return errors.Wrap(err, "failed to convert resource map to items")
		}
		return nil
	})

	cmd.Long = usage
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getDataFromFunctionConfig(fc interface{}) (interface{}, error) {
	f, ok := fc.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("function config %#v is not valid", fc)
	}
	return f["data"], nil
}
