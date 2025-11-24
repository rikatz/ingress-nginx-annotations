/*
Copyright 2018 The Kubernetes Authors.

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

package log

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	enableAccessLogAnnotation  = "enable-access-log"
	enableRewriteLogAnnotation = "enable-rewrite-log"
)

var logAnnotations = parser.Annotation{
	Group: "log",
	Annotations: parser.AnnotationFields{
		enableAccessLogAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This configuration setting allows you to control if this location should generate an access_log`,
		},
		enableRewriteLogAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This configuration setting allows you to control if this location should generate logs from the rewrite feature usage`,
		},
	},
}
