/*
Copyright 2019 The Kubernetes Authors.

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

package mirror

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	mirrorRequestBodyAnnotation = "mirror-request-body"
	mirrorTargetAnnotation      = "mirror-target"
	mirrorHostAnnotation        = "mirror-host"
)

var OnOffRegex = regexp.MustCompile(`^(on|off)$`)

var MirrorAnnotation = parser.Annotation{
	Group: "mirror",
	Annotations: parser.AnnotationFields{
		mirrorRequestBodyAnnotation: {
			Validator:     parser.ValidateRegex(OnOffRegex, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines if the request-body should be sent to the mirror backend. Can be 'on' or 'off'`,
		},
		mirrorTargetAnnotation: {
			Validator:     parser.ValidateServerName,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation enables a request to be mirrored to a mirror backend.`,
		},
		mirrorHostAnnotation: {
			Validator:     parser.ValidateServerName,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation defines if a specific Host header should be set for mirrored request.`,
		},
	},
}
