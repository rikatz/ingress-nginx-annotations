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

package ipdenylist

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	ipDenylistAnnotation = "denylist-source-range"
)

var denylistAnnotations = parser.Annotation{
	Group: "acl",
	Annotations: parser.AnnotationFields{
		ipDenylistAnnotation: {
			Validator:     parser.ValidateCIDRs,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium, // Failure on parsing this may cause undesired access
			Documentation: `This annotation allows setting a list of IPs and networks that should be blocked to access this Location`,
		},
	},
}
