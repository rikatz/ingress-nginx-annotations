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

package xforwardedprefix

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	xForwardedForPrefixAnnotation = "x-forwarded-prefix"
)

var xForwardedForAnnotations = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		xForwardedForPrefixAnnotation: {
			Validator: parser.ValidateRegex(parser.RegexPathWithCapture, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation can be used to add the non-standard X-Forwarded-Prefix header to the upstream request with a string value. It can 
			contain regular characters and captured groups specified as '$1', '$2', etc.`,
		},
	},
}
