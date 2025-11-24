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

package http2pushpreload

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	http2PushPreloadAnnotation = "http2-push-preload"
)

var http2PushPreloadAnnotations = parser.Annotation{
	Group: "http2",
	Annotations: parser.AnnotationFields{
		http2PushPreloadAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `Enables automatic conversion of preload links specified in the “Link” response header fields into push requests`,
		},
	},
}
