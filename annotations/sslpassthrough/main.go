/*
Copyright 2016 The Kubernetes Authors.

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

package sslpassthrough

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	sslPassthroughAnnotation = "ssl-passthrough"
)

var SSLPassthroughAnnotations = parser.Annotation{
	Group: "", // TBD
	Annotations: parser.AnnotationFields{
		sslPassthroughAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows regexes but on a very limited set
			Documentation: `This annotation instructs the controller to send TLS connections directly to the backend instead of letting NGINX decrypt the communication.`,
		},
	},
}
