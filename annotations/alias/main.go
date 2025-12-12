/*
Copyright 2017 The Kubernetes Authors.

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

package alias

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	serverAliasAnnotation = "server-alias"
)

var AliasAnnotation = parser.Annotation{
	Group: "alias",
	Annotations: parser.AnnotationFields{
		serverAliasAnnotation: {
			Validator: parser.ValidateArrayOfServerName,
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskHigh, // High as this allows regex chars
			Documentation: `this annotation can be used to define additional server 
			aliases for this Ingress`,
			GatewayAPI: `
Set additional hostnames on your .spec.hostnames.`,
		},
	},
}
