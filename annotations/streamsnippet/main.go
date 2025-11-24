/*
Copyright 2021 The Kubernetes Authors.

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

package streamsnippet

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	streamSnippetAnnotation = "stream-snippet"
)

var streamSnippetAnnotations = parser.Annotation{
	Group: "snippets",
	Annotations: parser.AnnotationFields{
		streamSnippetAnnotation: {
			Validator:     parser.ValidateNull,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskCritical, // Critical, this annotation is not validated at all and allows arbitrary configurations
			Documentation: `This annotation allows setting a custom NGINX configuration on a stream block. This annotation does not contain any validation and it's usage is not recommended!`,
		},
	},
}
