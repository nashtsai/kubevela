/*
Copyright 2021 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

const (
	// OutputFieldName is the reference of context base object
	OutputFieldName = "output"
	// OutputsFieldName is the reference of context Auxiliaries
	OutputsFieldName = "outputs"
	// ConfigFieldName is the reference of context config
	ConfigFieldName = "config"
	// ParameterFieldName is the keyword in CUE template to define users' input and the reference to the context parameter
	ParameterFieldName = "parameter"
	// ContextName is the name of context
	ContextName = "name"
	// ContextAppName is the appName of context
	ContextAppName = "appName"
	// ContextAppRevision is the revision name of app of context
	ContextAppRevision = "appRevision"
	// ContextAppRevisionNum is the revision num of app of context
	ContextAppRevisionNum = "appRevisionNum"
	// ContextNamespace is the namespace of the app
	ContextNamespace = "namespace"
	// OutputSecretName is used to store all secret names which are generated by cloud resource components
	OutputSecretName = "outputSecretName"
	// ContextCompRevisionName is the component revision name of context
	ContextCompRevisionName = "revision"
	// ContextComponents is the components of app
	ContextComponents = "components"
	// ContextComponentType is the component type of current trait binding with
	ContextComponentType = "componentType"
	// ComponentRevisionPlaceHolder is the component revision name placeHolder, this field will be replace with real value
	// after component be created
	ComponentRevisionPlaceHolder = "KUBEVELA_COMPONENT_REVISION_PLACEHOLDER"
	// ContextDataArtifacts is used to store unstructured resources of components
	ContextDataArtifacts = "artifacts"
)
