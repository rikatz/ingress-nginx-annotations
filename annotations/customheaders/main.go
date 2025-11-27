/*
Copyright 2023 The Kubernetes Authors.

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

package customheaders

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	customHeadersConfigMapAnnotation = "custom-headers"
)

var CustomHeadersAnnotation = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		customHeadersConfigMapAnnotation: {
			Validator: parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation sets the name of a ConfigMap that specifies headers to pass to the client.
			Only ConfigMaps on the same namespace are allowed`,
		},
	},
}
